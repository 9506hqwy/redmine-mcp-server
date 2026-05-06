package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type WatchersDestroyIssueRequest struct {
	IssueId int                                `json:"issue_id" jsonschema:"description=The ID of the issue."`
	UserId  int                                `json:"user_id" jsonschema:"description=The ID of the user."`
	Params  *client.WatchersDestroyIssueParams `json:"params,omitempty"`
}

func registerWatchersDestroyIssue(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WatchersDestroyIssueRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("watchers_destroy_issue",
		mcp.WithDescription("Deletes the watcher with the specified ID from the issue."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(watchersDestroyIssueHandler))
}

func watchersDestroyIssueHandler(ctx context.Context, request mcp.CallToolRequest, req WatchersDestroyIssueRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WatchersDestroyIssue(ctx, req.IssueId, req.UserId, req.Params, authorizationHeader))
}

type WatchersDestroyRequest struct {
	Params *client.WatchersDestroyParams `json:"params,omitempty"`
}

func registerWatchersDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WatchersDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("watchers_destroy",
		mcp.WithDescription("Deletes the watcher with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(watchersDestroyHandler))
}

func watchersDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req WatchersDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WatchersDestroy(ctx, req.Params, authorizationHeader))
}
