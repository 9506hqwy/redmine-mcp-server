package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type CustomFieldsIndexRequest struct {
	Params *client.CustomFieldsIndexParams `json:"params,omitempty"`
}

func registerCustomFieldsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&CustomFieldsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("custom_fields_index",
		mcp.WithDescription("Returns all custom field definitions."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(customFieldsIndexHandler))
}

func customFieldsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req CustomFieldsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.CustomFieldsIndex(ctx, req.Params, authorizationHeader))
}
