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

function write-parameter-definition() {
    local NAME="$1"
    local DESCRIPTION="$2"
    local EXAMPLES="$3"
    local IS_REQUIRED="$4"
    local ENUM_VALUES="$5"
    local DEFAULT_VALUE="$6"
    local TY="$7"

    DESCRIPTION=$(normalize "${DESCRIPTION}")
    DESCRIPTION="${DESCRIPTION//\\/\\\\}"
    DESCRIPTION="${DESCRIPTION//\"/\\\"}"

    if [[ $EXAMPLES != "null" && $EXAMPLES != "{}" ]]; then
        DESCRIPTION="${DESCRIPTION} (examples: ${EXAMPLES})"
    fi

    REQUIRED=""
    if [[ "${IS_REQUIRED}" == "true" ]]; then
        REQUIRED="mcp.Required(),"
    fi

    ENUM=""
    if [[ -n "${ENUM_VALUES}" && "${ENUM_VALUES}" != "\"\"" ]]; then
        ENUM="mcp.Enum(${ENUM_VALUES}),"
    fi

    DEFAULT=""

    if [[ "${TY}" == "integer" ]]; then
        if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
            # https://github.com/microsoft/vscode/issues/250599
            # https://github.com/microsoft/vscode-copilot-release/issues/13588
            #DEFAULT="mcp.DefaultNumber(float64(${DEFAULT_VALUE})),"
            DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
        fi

        cat >> "${FILE_PATH}" <<EOF
		mcp.WithNumber("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
    elif [[ "${TY}" == "boolean" ]]; then
        if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
            #DEFAULT="mcp.DefaultBool(${DEFAULT_VALUE}),"
            DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
        fi

        cat >> "${FILE_PATH}" <<EOF
		mcp.WithBoolean("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
    else
        if [[ "$DEFAULT_VALUE" != "null" && "$DEFAULT_VALUE" != "{}" ]]; then
            #DEFAULT="mcp.DefaultString(\"${DEFAULT_VALUE}\"),"
            DESCRIPTION="${DESCRIPTION} (default: ${DEFAULT_VALUE})"
        fi

        cat >> "${FILE_PATH}" <<EOF
		mcp.WithString("${NAME}",
			mcp.Description("${DESCRIPTION}"),
			${REQUIRED}
			${ENUM}
			${DEFAULT}
		),
EOF
    fi
}

function write-preceding() {
    local TAG="$1"

    FILE_PATH=$(filename "${TAG}")
    cat > "${FILE_PATH}" <<EOF
package redmine

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
	openapi_types "github.com/oapi-codegen/runtime/types"
)
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
	tool := mcp.NewTool("${TOOL_NAME}",
		mcp.WithDescription("${DESCRIPTION}"),
EOF

    # parameters
    IFS_ORG="${IFS}"
    IFS=$'\n' PARAMS=($(yq '.parameters[] | . style="flow"' <<<"${PATH_INFO}"))
    IFS="${IFS_ORG}"
    PARAMS+=("${PATH_PARAMS[@]}")
    index=0
    while [[ $index -lt ${#PARAMS[@]} ]]
    do
        PARAMETER="${PARAMS[$index]}"

        NAME=$(yq -r '.name' <<<"${PARAMETER}")
        DESCRIPTION=$(yq -r '.description' <<<"${PARAMETER}")
        EXAMPLES=$(yq -r '.examples' <<<"${PARAMETER}")
        IS_REQUIRED=$(yq -r '.required' <<<"${PARAMETER}")
        ENUM_VALUES=$(yq -r '[.schema.enum[] | "\"\(.)\""] | join(",")' <<<"${PARAMETER}")
        DEFAULT_VALUE=$(yq -r '.schema.default' <<<"${PARAMETER}")
        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")

        if [[ "${TY}" == "object" ]]; then
            yq ".schema.properties | to_entries[] | . style=\"flow\"" <<<"${PARAMETER}" | while read -r PROPERTY
            do
                PNAME=$(yq -r '.key' <<<"${PROPERTY}")
                DESCRIPTION=$(yq -r '.value.description' <<<"${PROPERTY}")
                EXAMPLES=$(yq -r '.value.examples' <<<"${PROPERTY}")
                IS_REQUIRED=$(yq -r '.value.required' <<<"${PROPERTY}")
                ENUM_VALUES=$(yq -r '[.value.enum[] | "\"\(.)\""] | join(",")' <<<"${PROPERTY}")
                DEFAULT_VALUE=$(yq -r '.value.default' <<<"${PROPERTY}")
                TY=$(yq -r '.value.type' <<<"${PROPERTY}")

                write-parameter-definition "${NAME}.${PNAME}" "${DESCRIPTION}" "${EXAMPLES}" "${IS_REQUIRED}" "${ENUM_VALUES}" "${DEFAULT_VALUE}" "${TY}"
            done
        else
            write-parameter-definition "${NAME}" "${DESCRIPTION}" "${EXAMPLES}" "${IS_REQUIRED}" "${ENUM_VALUES}" "${DEFAULT_VALUE}" "${TY}"
        fi

        index=$(( index + 1 ))
    done

    # complete
    cat >> "${FILE_PATH}" <<EOF
	)

	s.AddTool(tool, ${API_NAME,}Handler)
}
EOF
}

function write-handler-func() {
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

        PATH_ARG+=("${VAR_NAME}")

        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        if [[ "${TY}" == "integer" ]]; then
            PATH_STMT+=("${VAR_NAME} := request.GetInt(\"${NAME}\", math.MinInt)")
        elif [[ "${TY}" == "boolean" ]]; then
            PATH_STMT+=("${VAR_NAME} := request.GetBool(\"${NAME}\", false)")
        else
            PATH_STMT+=("${VAR_NAME} := request.GetString(\"${NAME}\", \"\")")
        fi
    done < <(echo "${OP_PATH}" | grep -oP '{[^}]+}')

    PATH_STMT_STR="$(IFS=$'\n'; echo "${PATH_STMT[*]}")"
    PATH_ARG_STR="$(IFS=,; echo "${PATH_ARG[*]}")"
    if [[ -n "${PATH_ARG_STR}" ]]; then
        PATH_ARG_STR=", ${PATH_ARG_STR}"
    fi

    PARSE_STMT="params := parse${API_NAME}(request)"
    PARAM_ARG=",&params"
    FILE_PATH=$(filename "$(yq -r '.tags[0]' <<<"${PATH_INFO}")")
    cat >> "${FILE_PATH}" <<EOF

func ${API_NAME,}Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	${PATH_STMT_STR}
	${PARSE_STMT}
	return toResult(c.${API_NAME}(ctx ${PATH_ARG_STR} ${PARAM_ARG}, authorizationHeader))
}
EOF
}

function write-parser-parameter() {
    local NAME="$1"
    local IS_REQUIRED="$2"
    local TY="$3"
    local FORMAT="$4"
    local DEFAULT_VALUE="$5"
    local ITEM_TY="$6"
    local PNAME="$7"

    VAR_NAME=${NAME//-/_} # '-' -> '_'
    VAR_NAME=${VAR_NAME//./_} # '.' -> '_'
    if [[ "${VAR_NAME}" == 'type' ]]; then
        VAR_NAME='_type'
    elif [[ "${VAR_NAME}" == 'params' ]]; then
        VAR_NAME='_params'
    fi

    REF_NAME="&${VAR_NAME}"
    if [[ "${IS_REQUIRED}" == "true" ]]; then
        REF_NAME="${VAR_NAME}"
    fi

    PARAM_NAME=$(capitalize "${VAR_NAME}")

    if [[ "${TY}" == "integer" ]]; then
        if [[ "$DEFAULT_VALUE" == "null" || "$DEFAULT_VALUE" == "{}" ]]; then
            unset DEFAULT_VALUE
        fi

        CONV_STMT=""

        cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetInt("${PNAME}", ${DEFAULT_VALUE:-math.MinInt})
	if ${VAR_NAME} != math.MinInt {
        ${CONV_STMT}
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
    elif [[ "${TY}" == "boolean" ]]; then
        if [[ "$DEFAULT_VALUE" == "null" || "$DEFAULT_VALUE" == "{}" ]]; then
            unset DEFAULT_VALUE
        fi

        cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetBool("${PNAME}", ${DEFAULT_VALUE:-false})
	params.${PARAM_NAME} = ${REF_NAME}
EOF
    elif [[ "${TY}" == "array" ]]; then
        if [[ "${ITEM_TY}" == "integer" ]]; then
            # TODO: error handling
            TY_NAME="int"
            CONV_STMT="intValue"
            cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${PNAME}", "")
	if ${VAR_NAME} != "" {
        ${VAR_NAME} := strings.Split(${VAR_NAME}, ",")
        var intSlice []${TY_NAME}
        for _, v := range ${VAR_NAME} {
            intValue, _ := strconv.Atoi(v)
            intSlice = append(intSlice, ${CONV_STMT})
        }
		params.${PARAM_NAME} = &intSlice
	}
EOF
        else
            cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${PNAME}", "")
	if ${VAR_NAME} != "" {
        ${VAR_NAME} := strings.Split(${VAR_NAME}, ",")
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
        fi
    elif [[ "${TY}" == "object" ]]; then
        METHOD_NAME="parse${PARAM_NAME}"
        if [[ "${NAME}" == "query" ]]; then
            METHOD_NAME="parse${API_NAME}${PARAM_NAME}"
            QUERY_METHOD_NAME="${METHOD_NAME}"
        fi

        cat >> "${FILE_PATH}" <<EOF

	params.${PARAM_NAME} = ${METHOD_NAME}(&request)
EOF
    else
        CONV_STMT=""
        if [[ "${FORMAT}" == "date-time" ]]; then
            # TODO: error handling
            CONV_STMT="${VAR_NAME}, _ := time.Parse(time.RFC3339, ${VAR_NAME})"
        elif [[ "${FORMAT}" == "date" ]]; then
            # TODO: error handling
            CONV_STMT="${VAR_NAME}, _ := time.Parse(time.DateOnly, ${VAR_NAME})"
            REF_NAME="&openapi_types.Date{Time: ${VAR_NAME}}"
        fi
        cat >> "${FILE_PATH}" <<EOF

	${VAR_NAME} := request.GetString("${PNAME}", "")
	if ${VAR_NAME} != "" {
        ${CONV_STMT}
		params.${PARAM_NAME} = ${REF_NAME}
	}
EOF
    fi
}

function write-parser-func() {
    local PATH_INFO="$1"
    local PATH_PARAMS="$2"

    TOOL_NAME=$(toolname "${PATH_INFO}")
    API_NAME=$(capitalize "${TOOL_NAME}")

    TAG=$(yq -r '.tags[0]' <<<"${PATH_INFO}")
    FILE_PATH=$(filename "${TAG}")
    QUERY_METHOD_NAME=''
    # signature
    cat >> "${FILE_PATH}" <<EOF

func parse${API_NAME}(request mcp.CallToolRequest) client.${API_NAME}Params {
	params := client.${API_NAME}Params{}
EOF

    # parameters
    IFS_ORG="${IFS}"
    IFS=$'\n' PARAMS=($(yq '.parameters[] | select(.in != "path") | . style="flow"' <<<"${PATH_INFO}"))
    IFS="${IFS_ORG}"
    PARAMS+=("${PATH_PARAMS[@]}")
    index=0
    while [[ $index -lt ${#PARAMS[@]} ]]
    do
        PARAMETER="${PARAMS[$index]}"

        NAME=$(yq -r '.name' <<<"${PARAMETER}")
        IS_REQUIRED=$(yq -r '.required' <<<"${PARAMETER}")
        TY=$(yq -r '.schema.type' <<<"${PARAMETER}")
        FORMAT=$(yq -r '.schema.format' <<<"${PARAMETER}")
        DEFAULT_VALUE=$(yq -r '.schema.default' <<<"${PARAMETER}")
        ITEM_TY=$(yq -r '.schema.items.type' <<<"${PARAMETER}")

        write-parser-parameter "${NAME}" "${IS_REQUIRED}" "${TY}" "${FORMAT}" "${DEFAULT_VALUE}" "${ITEM_TY}" "${NAME}"

        index=$(( index + 1 ))
    done

    # complete
    cat >> "${FILE_PATH}" <<EOF

	return params
}
EOF

    if [[ -z "${QUERY_METHOD_NAME}" ]]; then
        return
    fi

    if [[ "${TAG}" == "search" ]]; then
        return
    fi

    # write query func
    cat >> "${FILE_PATH}" <<EOF

func ${QUERY_METHOD_NAME}(request *mcp.CallToolRequest) *client.${API_NAME}Params_Query {
	params := client.${API_NAME}Params_Query{}
EOF

    yq '.parameters[] | select(.name == "query") | select(.schema.type == "object") | . style="flow"' <<<"${PATH_INFO}" | while read -r PARAMETER
    do
        yq ".schema.properties | to_entries[] | . style=\"flow\"" <<<"${PARAMETER}" | while read -r PROPERTY
        do
            PNAME=$(yq -r '.key' <<<"${PROPERTY}")
            IS_REQUIRED=$(yq -r '.value.required' <<<"${PROPERTY}")
            TY=$(yq -r '.value.type' <<<"${PROPERTY}")
            FORMAT=$(yq -r '.value.format' <<<"${PARAMETER}")
            DEFAULT_VALUE=$(yq -r '.value.default' <<<"${PROPERTY}")
            ITEM_TY=$(yq -r '.value.items.type' <<<"${PARAMETER}")

            write-parser-parameter "${PNAME}" "${IS_REQUIRED}" "${TY}" "${FORMAT}" "${DEFAULT_VALUE}" "${ITEM_TY}" "query.${PNAME}"
        done
    done

    # complete
    cat >> "${FILE_PATH}" <<EOF

	return &params
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
        write-register-func "${OP_DELETE}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_DELETE}"
        write-parser-func "${OP_DELETE}" "${PATH_PARAMS}"
    fi

    OP_POST=$(yq '.value.post | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.post.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_POST" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        write-register-func "${OP_POST}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_POST}"
        write-parser-func "${OP_POST}" "${PATH_PARAMS}"
    fi

    OP_PATCH=$(yq '.value.patch | . style="flow"' <<<"${PATH_INFO}")
    OP_REQ_BODY=$(yq '.value.patch.requestBody' <<<"${PATH_INFO}")
    if [[ "$OP_PATCH" != 'null' && "$OP_REQ_BODY" == 'null' ]]; then
        write-register-func "${OP_PATCH}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_PATCH}"
        write-parser-func "${OP_PATCH}" "${PATH_PARAMS}"
    fi

    OP_GET=$(yq '.value.get | . style="flow"' <<<"${PATH_INFO}")
    if [[ "$OP_GET" != 'null' ]]; then
        write-register-func "${OP_GET}" "${PATH_PARAMS}"
        write-handler-func "${OP_PATH}" "${OP_GET}"
        write-parser-func "${OP_GET}" "${PATH_PARAMS}"
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
