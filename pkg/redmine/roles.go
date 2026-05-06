package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type RolesIndexRequest struct {
	Params *client.RolesIndexParams `json:"params,omitempty"`
}

func registerRolesIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&RolesIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("roles_index",
		mcp.WithDescription("Returns a list of all roles."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rolesIndexHandler))
}

func rolesIndexHandler(ctx context.Context, request mcp.CallToolRequest, req RolesIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.RolesIndex(ctx, req.Params, authorizationHeader))
}

type RolesShowRequest struct {
	Id     int                     `json:"id" jsonschema:"description=The ID of the role."`
	Params *client.RolesShowParams `json:"params,omitempty"`
}

func registerRolesShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&RolesShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("roles_show",
		mcp.WithDescription("Returns the role with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(rolesShowHandler))
}

func rolesShowHandler(ctx context.Context, request mcp.CallToolRequest, req RolesShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.RolesShow(ctx, req.Id, req.Params, authorizationHeader))
}
