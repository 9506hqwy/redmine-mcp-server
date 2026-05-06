package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type SearchIndexProjectRequest struct {
	ProjectId string                           `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.SearchIndexProjectParams `json:"params,omitempty"`
}

func registerSearchIndexProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&SearchIndexProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("search_index_project",
		mcp.WithDescription("Returns search results based on the specified query parameters."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(searchIndexProjectHandler))
}

func searchIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest, req SearchIndexProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.SearchIndexProject(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type SearchIndexRequest struct {
	Params *client.SearchIndexParams `json:"params,omitempty"`
}

func registerSearchIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&SearchIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("search_index",
		mcp.WithDescription("Returns search results based on the specified query parameters."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(searchIndexHandler))
}

func searchIndexHandler(ctx context.Context, request mcp.CallToolRequest, req SearchIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.SearchIndex(ctx, req.Params, authorizationHeader))
}
