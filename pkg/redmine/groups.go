package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type GroupsIndexRequest struct {
	Params *client.GroupsIndexParams `json:"params,omitempty"`
}

func registerGroupsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GroupsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("groups_index",
		mcp.WithDescription("Returns a list of all groups."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(groupsIndexHandler))
}

func groupsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req GroupsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GroupsIndex(ctx, req.Params, authorizationHeader))
}

type GroupsDestroyRequest struct {
	Id     int                         `json:"id" jsonschema:"description=The ID of the group."`
	Params *client.GroupsDestroyParams `json:"params,omitempty"`
}

func registerGroupsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GroupsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("groups_destroy",
		mcp.WithDescription("Deletes the group with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(groupsDestroyHandler))
}

func groupsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req GroupsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GroupsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type GroupsShowRequest struct {
	Id     int                      `json:"id" jsonschema:"description=The ID of the group."`
	Params *client.GroupsShowParams `json:"params,omitempty"`
}

func registerGroupsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GroupsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("groups_show",
		mcp.WithDescription("Returns the group with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(groupsShowHandler))
}

func groupsShowHandler(ctx context.Context, request mcp.CallToolRequest, req GroupsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GroupsShow(ctx, req.Id, req.Params, authorizationHeader))
}

type GroupsRemoveUserRequest struct {
	Id     int                            `json:"id" jsonschema:"description=The ID of the group."`
	UserId int                            `json:"user_id" jsonschema:"description=The ID of the user."`
	Params *client.GroupsRemoveUserParams `json:"params,omitempty"`
}

func registerGroupsRemoveUser(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GroupsRemoveUserRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("groups_remove_user",
		mcp.WithDescription("Removes a user from a group."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(groupsRemoveUserHandler))
}

func groupsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest, req GroupsRemoveUserRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GroupsRemoveUser(ctx, req.Id, req.UserId, req.Params, authorizationHeader))
}
