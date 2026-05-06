package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type ProjectsIndexCsvRequest struct {
	Params *client.ProjectsIndexCsvParams `json:"params,omitempty"`
}

func registerProjectsIndexCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsIndexCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_index_csv",
		mcp.WithDescription("Returns all projects (including all public projects and private projects to which the user has access)."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsIndexCsvHandler))
}

func projectsIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsIndexCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsIndexCsv(ctx, req.Params, authorizationHeader))
}

type ProjectsCreateRequest struct {
	Params *client.ProjectsCreateParams         `json:"params,omitempty"`
	Body   client.ProjectsCreateJSONRequestBody `json:"body,omitempty"`
}

func registerProjectsCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsCreateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_create",
		mcp.WithDescription("Creates a new project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsCreateHandler))
}

func projectsCreateHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsCreateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsCreate(ctx, req.Params, req.Body, authorizationHeader))
}

type ProjectsIndexRequest struct {
	Params *client.ProjectsIndexParams `json:"params,omitempty"`
}

func registerProjectsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_index",
		mcp.WithDescription("Returns all projects (including all public projects and private projects to which the user has access)."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsIndexHandler))
}

func projectsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsIndex(ctx, req.Params, authorizationHeader))
}

type ProjectsDestroyRequest struct {
	Id     string                        `json:"id" jsonschema:"description=The ID or identifier of the project."`
	Params *client.ProjectsDestroyParams `json:"params,omitempty"`
}

func registerProjectsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_destroy",
		mcp.WithDescription("Deletes the project with the specified ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsDestroyHandler))
}

func projectsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type ProjectsUpdatePatchRequest struct {
	Id     string                                    `json:"id" jsonschema:"description=The ID or identifier of the project."`
	Params *client.ProjectsUpdatePatchParams         `json:"params,omitempty"`
	Body   client.ProjectsUpdatePatchJSONRequestBody `json:"body,omitempty"`
}

func registerProjectsUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_update_patch",
		mcp.WithDescription("Updates the project with the specified ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsUpdatePatchHandler))
}

func projectsUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}

type ProjectsShowRequest struct {
	Id     string                     `json:"id" jsonschema:"description=The ID or identifier of the project."`
	Params *client.ProjectsShowParams `json:"params,omitempty"`
}

func registerProjectsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_show",
		mcp.WithDescription("Returns the project with the specified ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsShowHandler))
}

func projectsShowHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsShow(ctx, req.Id, req.Params, authorizationHeader))
}

type ProjectsArchivePostRequest struct {
	Id     string                            `json:"id" jsonschema:"description=The ID or identifier of the project."`
	Params *client.ProjectsArchivePostParams `json:"params,omitempty"`
}

func registerProjectsArchivePost(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsArchivePostRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_archive_post",
		mcp.WithDescription("Archives the project with the specified ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsArchivePostHandler))
}

func projectsArchivePostHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsArchivePostRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsArchivePost(ctx, req.Id, req.Params, authorizationHeader))
}

type ProjectsUnarchivePostRequest struct {
	Id     string                              `json:"id" jsonschema:"description=The ID or identifier of the project."`
	Params *client.ProjectsUnarchivePostParams `json:"params,omitempty"`
}

func registerProjectsUnarchivePost(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&ProjectsUnarchivePostRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("projects_unarchive_post",
		mcp.WithDescription("Unarchives the project with the specified ID or identifier."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(projectsUnarchivePostHandler))
}

func projectsUnarchivePostHandler(ctx context.Context, request mcp.CallToolRequest, req ProjectsUnarchivePostRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.ProjectsUnarchivePost(ctx, req.Id, req.Params, authorizationHeader))
}
