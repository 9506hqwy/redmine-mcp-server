package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type IssueRelationsCreateRequest struct {
	IssueId int                                        `json:"issue_id" jsonschema:"description=The ID of the issue."`
	Params  *client.IssueRelationsCreateParams         `json:"params,omitempty"`
	Body    client.IssueRelationsCreateJSONRequestBody `json:"body"`
}

func registerIssueRelationsCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueRelationsCreateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_relations_create",
		mcp.WithDescription("Creates a relation for the specified issue ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueRelationsCreateHandler))
}

func issueRelationsCreateHandler(ctx context.Context, request mcp.CallToolRequest, req IssueRelationsCreateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueRelationsCreate(ctx, req.IssueId, req.Params, req.Body, authorizationHeader))
}

type IssueRelationsIndexRequest struct {
	IssueId int                               `json:"issue_id" jsonschema:"description=The ID of the issue."`
	Params  *client.IssueRelationsIndexParams `json:"params,omitempty"`
}

func registerIssueRelationsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueRelationsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_relations_index",
		mcp.WithDescription("Returns the relations for the specified issue ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueRelationsIndexHandler))
}

func issueRelationsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req IssueRelationsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueRelationsIndex(ctx, req.IssueId, req.Params, authorizationHeader))
}

type IssueRelationsDestroyRequest struct {
	Id     int                                 `json:"id" jsonschema:"description=The ID of the relation."`
	Params *client.IssueRelationsDestroyParams `json:"params,omitempty"`
}

func registerIssueRelationsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueRelationsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_relations_destroy",
		mcp.WithDescription("Deletes the relation with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueRelationsDestroyHandler))
}

func issueRelationsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req IssueRelationsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueRelationsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type IssueRelationsShowRequest struct {
	Id     int                              `json:"id" jsonschema:"description=The ID of the relation."`
	Params *client.IssueRelationsShowParams `json:"params,omitempty"`
}

func registerIssueRelationsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueRelationsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_relations_show",
		mcp.WithDescription("Returns the relation with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueRelationsShowHandler))
}

func issueRelationsShowHandler(ctx context.Context, request mcp.CallToolRequest, req IssueRelationsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueRelationsShow(ctx, req.Id, req.Params, authorizationHeader))
}
