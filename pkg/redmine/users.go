package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type UsersIndexCsvRequest struct {
	Params *client.UsersIndexCsvParams `json:"params,omitempty"`
}

func registerUsersIndexCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&UsersIndexCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("users_index_csv",
		mcp.WithDescription("Returns a list of all users in CSV format."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersIndexCsvHandler))
}

func usersIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest, req UsersIndexCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.UsersIndexCsv(ctx, req.Params, authorizationHeader))
}

type UsersIndexRequest struct {
	Params *client.UsersIndexParams `json:"params,omitempty"`
}

func registerUsersIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&UsersIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("users_index",
		mcp.WithDescription("Returns a list of all users."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersIndexHandler))
}

func usersIndexHandler(ctx context.Context, request mcp.CallToolRequest, req UsersIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.UsersIndex(ctx, req.Params, authorizationHeader))
}

type UsersDestroyRequest struct {
	Id     string                     `json:"id" jsonschema:"description=The ID or 'current' of the user."`
	Params *client.UsersDestroyParams `json:"params,omitempty"`
}

func registerUsersDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&UsersDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("users_destroy",
		mcp.WithDescription("Deletes the user with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersDestroyHandler))
}

func usersDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req UsersDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.UsersDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type UsersShowRequest struct {
	Id     string                  `json:"id" jsonschema:"description=The ID or 'current' of the user."`
	Params *client.UsersShowParams `json:"params,omitempty"`
}

func registerUsersShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&UsersShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("users_show",
		mcp.WithDescription("Returns the user with the specified ID. Use /users/current.json to retrieve the user whose credentials is used to access the API."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(usersShowHandler))
}

func usersShowHandler(ctx context.Context, request mcp.CallToolRequest, req UsersShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.UsersShow(ctx, req.Id, req.Params, authorizationHeader))
}
