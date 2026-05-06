#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname "$(dirname "$0")")
TMP_DIR=/tmp
BASE_DIR="${BASE_DIR}/pkg/redmine"

curl -fsSL \
    -o "${TMP_DIR}/openapi.yml" \
    https://github.com/9506hqwy/openapi-spec-redmine/raw/refs/heads/main/openapi.yml

function capitalize() {
    VALUE="$1"
    # shellcheck disable=SC2206
    IFS=_ ARR=(${VALUE})

    CAP_ARR=()
    for V in "${ARR[@]}"
    do
        CAP_ARR+=("${V^}")
    done

    IFS='' CAP="${CAP_ARR[*]}"

    echo "${CAP}"
}

function filename() {
    TAG="$1"
    echo "${BASE_DIR}/${TAG}.go"
}

function toolname() {
    PATH_INFO="$1"
    OP_ID=$(yq -r '.operationId' <<<"${PATH_INFO}")
    TOO_NAME=${OP_ID//#/_}    # '#' -> '_'
    TOO_NAME=${TOO_NAME//-/_} # '-' -> '_'
    echo "${TOO_NAME}"
}

function normalize() {
    VALUE="$1"
    # shellcheck disable=SC2206
    ARR=(${VALUE})
    echo "${ARR[*]}"
}

function write-preceding() {
    local TAG="$1"

    FILE_PATH=$(filename "${TAG}")
    cat > "${FILE_PATH}" <<EOF
package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)
EOF
}

function write-request-struct(){
    local OP_PATH="$1"
    local PATH_INFO="$2"

    TOOL_NAME=$(toolname "${PATH_INFO}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    # parameters
    PATH_STMT=()
    PATH_ARG=()
    while read -r NAME
    do
        NAME="${NAME//{/}" # '{' -> ''
        NAME="${NAME//\}/}" # '}' -> ''

        PARAMETER=$(yq ".parameters[] | select(.name == \"${NAME}\")" <<<"${PATH_INFO}")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        PARAM_NAME=$(capitalize "${VAR_NAME}")

        PATH_ARG+=("${PARAM_NAME}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        DESCRIPTION=$(yq -r '.description' <<<"${PARAMETER}")

        DESCRIPTION=$(normalize "${DESCRIPTION}")
        DESCRIPTION="${DESCRIPTION//\`/\'}" # '`' -> '\''

        if [[ "${TY}" == "integer" ]]; then
            PATH_STMT+=("${PARAM_NAME} int \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
        elif [[ "${TY}" == "boolean" ]]; then
            PATH_STMT+=("${PARAM_NAME} bool \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
        else
            PATH_STMT+=("${PARAM_NAME} string \`json:\"${NAME}\" jsonschema:\"description=${DESCRIPTION}\"\`")
        fi
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_STMT_STR="$(IFS=$'\n'; echo "${PATH_STMT[*]}")"

    FILE_PATH=$(filename "$(yq -r '.tags[0]' <<<"${PATH_INFO}")")
    cat >> "${FILE_PATH}" <<EOF
type ${API_NAME}Request struct {
    ${PATH_STMT_STR}
    Params *client.${API_NAME}Params \`json:"params,omitempty"\`
}
EOF

}

function write-register-func() {
    local PATH_INFO="$1"
    local PATH_PARAMS="$2"

    DESCRIPTION=$(yq -r '.description' <<<"${PATH_INFO}")
    DESCRIPTION=$(normalize "${DESCRIPTION}")
    DESCRIPTION="${DESCRIPTION//\"/\\\"}"

    TOOL_NAME=$(toolname "${PATH_INFO}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    echo "| ${TOOL_NAME} | ${DESCRIPTION} |"

    FILE_PATH=$(filename "$(yq -r '.tags[0]' <<<"${PATH_INFO}")")
    # signature
    cat >> "${FILE_PATH}" <<EOF

func register${API_NAME}(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&${API_NAME}Request{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("${TOOL_NAME}",
		mcp.WithDescription("${DESCRIPTION}"),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(${API_NAME,}Handler))
}
EOF
}

function write-handler-func() {
    local OP_PATH="$1"
    local PATH_INFO="$2"

    TOOL_NAME=$(toolname "${PATH_INFO}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    # parameters
    PATH_ARG=()
    while read -r NAME
    do
        NAME="${NAME//{/}" # '{' -> ''
        NAME="${NAME//\}/}" # '}' -> ''

        PARAMETER=$(yq ".parameters[] | select(.name == \"${NAME}\")" <<<"${PATH_INFO}")

        VAR_NAME=${NAME//-/_} # '-' -> '_'
        if [[ "${VAR_NAME}" == 'type' ]]; then
            VAR_NAME='_type'
        elif [[ "${VAR_NAME}" == 'params' ]]; then
            VAR_NAME='_params'
        fi

        PARAM_NAME=$(capitalize "${VAR_NAME}")

        PATH_ARG+=("req.${PARAM_NAME}")
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_ARG_STR="$(IFS=,; echo "${PATH_ARG[*]}")"
    if [[ -n "${PATH_ARG_STR}" ]]; then
        PATH_ARG_STR=", ${PATH_ARG_STR}"
    fi

    FILE_PATH=$(filename "$(yq -r '.tags[0]' <<<"${PATH_INFO}")")
    cat >> "${FILE_PATH}" <<EOF

func ${API_NAME,}Handler(ctx context.Context, request mcp.CallToolRequest, req ${API_NAME}Request) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.${API_NAME}(ctx ${PATH_ARG_STR}, req.Params, authorizationHeader))
}
EOF
}

yq -r '.paths | .. | select(has("tags")) | .tags | .[]' "${TMP_DIR}/openapi.yml" | while read -r TAG
do
    FILE_PATH=$(filename "${TAG}")
    rm -rf "${FILE_PATH}"
done

yq -r '.paths | .. | select(has("tags")) | .tags | .[]' "${TMP_DIR}/openapi.yml" | while read -r TAG
do
    write-preceding "${TAG}"
done

yq '.paths | to_entries | .[] | . style="flow"' "${TMP_DIR}/openapi.yml" | while read -r PATH_INFO
do
    OP_PATH=$(yq -r '.key' <<<"${PATH_INFO}")
    PATH_PARAMS=$(yq '.value.parameters[] | . style="flow"' <<<"${PATH_INFO}")

    OP_DELETE=$(yq '.value.delete | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_DELETE" != 'null' ]]; then
        write-request-struct "${OP_PATH}" "${OP_DELETE}"
        write-register-func "${OP_DELETE}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_DELETE}"
    fi

    OP_POST=$(yq '.value.post | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.post.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_POST" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        write-request-struct "${OP_PATH}" "${OP_POST}"
        write-register-func "${OP_POST}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_POST}"
    fi

    OP_PATCH=$(yq '.value.patch | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.patch.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_PATCH" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        write-request-struct "${OP_PATH}" "${OP_PATCH}"
        write-register-func "${OP_PATCH}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_PATCH}"
    fi

    OP_GET=$(yq '.value.get | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_GET" != 'null' ]]; then
        write-request-struct "${OP_PATH}" "${OP_GET}"
        write-register-func "${OP_GET}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_GET}"
    fi
done

# write tools.go
TOOLS_PATH="${BASE_DIR}/tools.go"
cat > "${TOOLS_PATH}" <<EOF
package redmine

import (
	"github.com/mark3labs/mcp-go/server"
)

func RegisterTools(s *server.MCPServer, readonly bool) {
EOF

yq '.paths | to_entries | .[] | . style="flow"' "${TMP_DIR}/openapi.yml" | while read -r PATH_INFO
do
    OP_PATH=$(yq -r '.key' <<<"${PATH_INFO}")

    # If the tool name is too long, we need to use a shorter name
    # to avoid exceeding the maximum length of 46 characters.
    # See https://github.com/microsoft/vscode/blob/1.101.2/src/vs/workbench/contrib/mcp/common/mcpTypes.ts#L710-L714

    OP_DELETE=$(yq '.value.delete | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_DELETE" != 'null' ]]; then
        TOOL_NAME=$(toolname "${OP_DELETE}")
        API_NAME=$(capitalize "${TOOL_NAME}")

        if [[ ${#TOOL_NAME} -gt 46 ]]; then
            echo "//if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_POST=$(yq '.value.post | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.post.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_POST" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        TOOL_NAME=$(toolname "${OP_POST}")
        API_NAME=$(capitalize "${TOOL_NAME}")

        if [[ ${#TOOL_NAME} -gt 46 ]]; then
            echo "//if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_PATCH=$(yq '.value.patch | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.patch.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_PATCH" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        TOOL_NAME=$(toolname "${OP_PATCH}")
        API_NAME=$(capitalize "${TOOL_NAME}")

        if [[ ${#TOOL_NAME} -gt 46 ]]; then
            echo "//if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        else
            echo "if !readonly { register${API_NAME}(s) }" >> "${TOOLS_PATH}"
        fi
    fi

    OP_GET=$(yq '.value.get | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_GET" != 'null' ]]; then
        TOOL_NAME=$(toolname "${OP_GET}")
        API_NAME=$(capitalize "${TOOL_NAME}")

        if [[ ${#TOOL_NAME} -gt 46 ]]; then
            echo "//register${API_NAME}(s)" >> "${TOOLS_PATH}"
        else
            echo "register${API_NAME}(s)" >> "${TOOLS_PATH}"
        fi
    fi
done

cat >> "${TOOLS_PATH}" <<EOF
}
EOF

# format
go fmt ./...
