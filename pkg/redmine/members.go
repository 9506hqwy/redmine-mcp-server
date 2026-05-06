package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type MembersDestroyRequest struct {
	Id     int                          `json:"id" jsonschema:"description=The ID of the membership."`
	Params *client.MembersDestroyParams `json:"params,omitempty"`
}

func registerMembersDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&MembersDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("members_destroy",
		mcp.WithDescription("Deletes the membership with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(membersDestroyHandler))
}

func membersDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req MembersDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.MembersDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type MembersShowRequest struct {
	Id     int                       `json:"id" jsonschema:"description=The ID of the membership."`
	Params *client.MembersShowParams `json:"params,omitempty"`
}

func registerMembersShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&MembersShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("members_show",
		mcp.WithDescription("Returns the membership with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(membersShowHandler))
}

func membersShowHandler(ctx context.Context, request mcp.CallToolRequest, req MembersShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.MembersShow(ctx, req.Id, req.Params, authorizationHeader))
}

type MembersIndexRequest struct {
	ProjectId string                     `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.MembersIndexParams `json:"params,omitempty"`
}

func registerMembersIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&MembersIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("members_index",
		mcp.WithDescription("Returns a paginated list of project memberships."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(membersIndexHandler))
}

func membersIndexHandler(ctx context.Context, request mcp.CallToolRequest, req MembersIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.MembersIndex(ctx, req.ProjectId, req.Params, authorizationHeader))
}
