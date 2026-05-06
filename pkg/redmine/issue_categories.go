package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type IssueCategoriesDestroyRequest struct {
	Id     int                                  `json:"id" jsonschema:"description=The ID of the issue category."`
	Params *client.IssueCategoriesDestroyParams `json:"params,omitempty"`
}

func registerIssueCategoriesDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueCategoriesDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_categories_destroy",
		mcp.WithDescription("Deletes the issue category with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueCategoriesDestroyHandler))
}

func issueCategoriesDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req IssueCategoriesDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueCategoriesDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type IssueCategoriesShowRequest struct {
	Id     int                               `json:"id" jsonschema:"description=The ID of the issue category."`
	Params *client.IssueCategoriesShowParams `json:"params,omitempty"`
}

func registerIssueCategoriesShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueCategoriesShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_categories_show",
		mcp.WithDescription("Returns the issue category with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueCategoriesShowHandler))
}

func issueCategoriesShowHandler(ctx context.Context, request mcp.CallToolRequest, req IssueCategoriesShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueCategoriesShow(ctx, req.Id, req.Params, authorizationHeader))
}

type IssueCategoriesIndexRequest struct {
	ProjectId string                             `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.IssueCategoriesIndexParams `json:"params,omitempty"`
}

func registerIssueCategoriesIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueCategoriesIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_categories_index",
		mcp.WithDescription("Returns the issue categories available for the specified project by ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueCategoriesIndexHandler))
}

func issueCategoriesIndexHandler(ctx context.Context, request mcp.CallToolRequest, req IssueCategoriesIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueCategoriesIndex(ctx, req.ProjectId, req.Params, authorizationHeader))
}
