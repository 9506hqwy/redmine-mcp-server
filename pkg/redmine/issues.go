package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type IssuesIndexCsvRequest struct {
	Params *client.IssuesIndexCsvParams `json:"params,omitempty"`
}

func registerIssuesIndexCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index_csv",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexCsvHandler))
}

func issuesIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndexCsv(ctx, req.Params, authorizationHeader))
}

type IssuesCreateRequest struct {
	Params *client.IssuesCreateParams         `json:"params,omitempty"`
	Body   client.IssuesCreateJSONRequestBody `json:"body,omitempty"`
}

func registerIssuesCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesCreateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_create",
		mcp.WithDescription("Creates a new issue."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesCreateHandler))
}

func issuesCreateHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesCreateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesCreate(ctx, req.Params, req.Body, authorizationHeader))
}

type IssuesIndexRequest struct {
	Params *client.IssuesIndexParams `json:"params,omitempty"`
}

func registerIssuesIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexHandler))
}

func issuesIndexHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndex(ctx, req.Params, authorizationHeader))
}

type IssuesIndexPdfRequest struct {
	Params *client.IssuesIndexPdfParams `json:"params,omitempty"`
}

func registerIssuesIndexPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index_pdf",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexPdfHandler))
}

func issuesIndexPdfHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndexPdf(ctx, req.Params, authorizationHeader))
}

type IssuesDestroyRequest struct {
	Id     int                         `json:"id" jsonschema:"description=The ID of the issue."`
	Params *client.IssuesDestroyParams `json:"params,omitempty"`
}

func registerIssuesDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_destroy",
		mcp.WithDescription("Deletes the issue with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesDestroyHandler))
}

func issuesDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type IssuesUpdatePatchRequest struct {
	Id     int                                     `json:"id" jsonschema:"description=The ID of the issue."`
	Params *client.IssuesUpdatePatchParams         `json:"params,omitempty"`
	Body   client.IssuesUpdatePatchJSONRequestBody `json:"body,omitempty"`
}

func registerIssuesUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_update_patch",
		mcp.WithDescription("Updates the issue with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesUpdatePatchHandler))
}

func issuesUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}

type IssuesShowRequest struct {
	Id     int                      `json:"id" jsonschema:"description=The ID of the issue."`
	Params *client.IssuesShowParams `json:"params,omitempty"`
}

func registerIssuesShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_show",
		mcp.WithDescription("Returns the issue with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesShowHandler))
}

func issuesShowHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesShow(ctx, req.Id, req.Params, authorizationHeader))
}

type IssuesShowPdfRequest struct {
	Id     int                         `json:"id" jsonschema:"description=The ID of the issue."`
	Params *client.IssuesShowPdfParams `json:"params,omitempty"`
}

func registerIssuesShowPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesShowPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_show_pdf",
		mcp.WithDescription("Returns the issue with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesShowPdfHandler))
}

func issuesShowPdfHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesShowPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesShowPdf(ctx, req.Id, req.Params, authorizationHeader))
}

type IssuesIndexProjectCsvRequest struct {
	ProjectId string                              `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.IssuesIndexProjectCsvParams `json:"params,omitempty"`
}

func registerIssuesIndexProjectCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexProjectCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index_project_csv",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexProjectCsvHandler))
}

func issuesIndexProjectCsvHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexProjectCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndexProjectCsv(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type IssuesCreateProjectRequest struct {
	ProjectId string                                    `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.IssuesCreateProjectParams         `json:"params,omitempty"`
	Body      client.IssuesCreateProjectJSONRequestBody `json:"body,omitempty"`
}

func registerIssuesCreateProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesCreateProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_create_project",
		mcp.WithDescription("Creates a new issue."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesCreateProjectHandler))
}

func issuesCreateProjectHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesCreateProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesCreateProject(ctx, req.ProjectId, req.Params, req.Body, authorizationHeader))
}

type IssuesIndexProjectRequest struct {
	ProjectId string                           `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.IssuesIndexProjectParams `json:"params,omitempty"`
}

func registerIssuesIndexProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index_project",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexProjectHandler))
}

func issuesIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndexProject(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type IssuesIndexProjectPdfRequest struct {
	ProjectId string                              `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.IssuesIndexProjectPdfParams `json:"params,omitempty"`
}

func registerIssuesIndexProjectPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssuesIndexProjectPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issues_index_project_pdf",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issuesIndexProjectPdfHandler))
}

func issuesIndexProjectPdfHandler(ctx context.Context, request mcp.CallToolRequest, req IssuesIndexProjectPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssuesIndexProjectPdf(ctx, req.ProjectId, req.Params, authorizationHeader))
}
