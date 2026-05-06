package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type EnumerationsIndexDocumentCategoryRequest struct {
	Params *client.EnumerationsIndexDocumentCategoryParams `json:"params,omitempty"`
}

func registerEnumerationsIndexDocumentCategory(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&EnumerationsIndexDocumentCategoryRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("enumerations_index_document_category",
		mcp.WithDescription("Returns a list of document categories."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(enumerationsIndexDocumentCategoryHandler))
}

func enumerationsIndexDocumentCategoryHandler(ctx context.Context, request mcp.CallToolRequest, req EnumerationsIndexDocumentCategoryRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.EnumerationsIndexDocumentCategory(ctx, req.Params, authorizationHeader))
}

type EnumerationsIndexIssuePriorityRequest struct {
	Params *client.EnumerationsIndexIssuePriorityParams `json:"params,omitempty"`
}

func registerEnumerationsIndexIssuePriority(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&EnumerationsIndexIssuePriorityRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("enumerations_index_issue_priority",
		mcp.WithDescription("Returns a list of issue priorities."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(enumerationsIndexIssuePriorityHandler))
}

func enumerationsIndexIssuePriorityHandler(ctx context.Context, request mcp.CallToolRequest, req EnumerationsIndexIssuePriorityRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.EnumerationsIndexIssuePriority(ctx, req.Params, authorizationHeader))
}

type EnumerationsIndexTimeEntryActivityRequest struct {
	Params *client.EnumerationsIndexTimeEntryActivityParams `json:"params,omitempty"`
}

func registerEnumerationsIndexTimeEntryActivity(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&EnumerationsIndexTimeEntryActivityRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("enumerations_index_time_entry_activity",
		mcp.WithDescription("Returns a list of time entry activities."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(enumerationsIndexTimeEntryActivityHandler))
}

func enumerationsIndexTimeEntryActivityHandler(ctx context.Context, request mcp.CallToolRequest, req EnumerationsIndexTimeEntryActivityRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.EnumerationsIndexTimeEntryActivity(ctx, req.Params, authorizationHeader))
}
