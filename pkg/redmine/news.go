package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type NewsCreateRequest struct {
	Params *client.NewsCreateParams         `json:"params,omitempty"`
	Body   client.NewsCreateJSONRequestBody `json:"body"`
}

func registerNewsCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsCreateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_create",
		mcp.WithDescription("Creates a new news item."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsCreateHandler))
}

func newsCreateHandler(ctx context.Context, request mcp.CallToolRequest, req NewsCreateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsCreate(ctx, req.Params, req.Body, authorizationHeader))
}

type NewsIndexRequest struct {
	Params *client.NewsIndexParams `json:"params,omitempty"`
}

func registerNewsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_index",
		mcp.WithDescription("Returns all news items across all projects with pagination."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsIndexHandler))
}

func newsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req NewsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsIndex(ctx, req.Params, authorizationHeader))
}

type NewsDestroyRequest struct {
	Id     int                       `json:"id" jsonschema:"description=The ID of the news."`
	Params *client.NewsDestroyParams `json:"params,omitempty"`
}

func registerNewsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_destroy",
		mcp.WithDescription("Deletes the news with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsDestroyHandler))
}

func newsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req NewsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type NewsUpdatePatchRequest struct {
	Id     int                                   `json:"id" jsonschema:"description=The ID of the news."`
	Params *client.NewsUpdatePatchParams         `json:"params,omitempty"`
	Body   client.NewsUpdatePatchJSONRequestBody `json:"body"`
}

func registerNewsUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_update_patch",
		mcp.WithDescription("Updates the news with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsUpdatePatchHandler))
}

func newsUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req NewsUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}

type NewsShowRequest struct {
	Id     int                    `json:"id" jsonschema:"description=The ID of the news."`
	Params *client.NewsShowParams `json:"params,omitempty"`
}

func registerNewsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_show",
		mcp.WithDescription("Returns the news item with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsShowHandler))
}

func newsShowHandler(ctx context.Context, request mcp.CallToolRequest, req NewsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsShow(ctx, req.Id, req.Params, authorizationHeader))
}

type NewsCreateProjectRequest struct {
	ProjectId string                                  `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.NewsCreateProjectParams         `json:"params,omitempty"`
	Body      client.NewsCreateProjectJSONRequestBody `json:"body"`
}

func registerNewsCreateProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsCreateProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_create_project",
		mcp.WithDescription("Creates a new news item."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsCreateProjectHandler))
}

func newsCreateProjectHandler(ctx context.Context, request mcp.CallToolRequest, req NewsCreateProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsCreateProject(ctx, req.ProjectId, req.Params, req.Body, authorizationHeader))
}

type NewsIndexProjectRequest struct {
	ProjectId string                         `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.NewsIndexProjectParams `json:"params,omitempty"`
}

func registerNewsIndexProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&NewsIndexProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("news_index_project",
		mcp.WithDescription("Returns all news items across all projects with pagination."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(newsIndexProjectHandler))
}

func newsIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest, req NewsIndexProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.NewsIndexProject(ctx, req.ProjectId, req.Params, authorizationHeader))
}
