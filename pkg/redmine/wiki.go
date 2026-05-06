package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type WikiShowRootRequest struct {
	ProjectId string                     `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.WikiShowRootParams `json:"params,omitempty"`
}

func registerWikiShowRoot(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowRootRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_root",
		mcp.WithDescription("Returns the details of the root wiki page."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowRootHandler))
}

func wikiShowRootHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowRootRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowRoot(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type WikiIndexRequest struct {
	ProjectId string                  `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.WikiIndexParams `json:"params,omitempty"`
}

func registerWikiIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_index",
		mcp.WithDescription("Returns a list of all pages in the project wiki."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiIndexHandler))
}

func wikiIndexHandler(ctx context.Context, request mcp.CallToolRequest, req WikiIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiIndex(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type WikiDestroyRequest struct {
	ProjectId string                    `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                    `json:"id" jsonschema:"description=The title of the wiki."`
	Params    *client.WikiDestroyParams `json:"params,omitempty"`
}

func registerWikiDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_destroy",
		mcp.WithDescription("Deletes a wiki page, its attachments and its history with the specified ID. If the deleted page is a parent page, its child pages are not deleted but changed as root pages."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiDestroyHandler))
}

func wikiDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req WikiDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiDestroy(ctx, req.ProjectId, req.Id, req.Params, authorizationHeader))
}

type WikiUpdatePatchRequest struct {
	ProjectId string                                `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                                `json:"id" jsonschema:"description=The title of the wiki."`
	Params    *client.WikiUpdatePatchParams         `json:"params,omitempty"`
	Body      client.WikiUpdatePatchJSONRequestBody `json:"body,omitempty"`
}

func registerWikiUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_update_patch",
		mcp.WithDescription("Creates or updates a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiUpdatePatchHandler))
}

func wikiUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req WikiUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiUpdatePatch(ctx, req.ProjectId, req.Id, req.Params, req.Body, authorizationHeader))
}

type WikiShowRequest struct {
	ProjectId string                 `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                 `json:"id" jsonschema:"description=The title of the wiki."`
	Params    *client.WikiShowParams `json:"params,omitempty"`
}

func registerWikiShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowHandler))
}

func wikiShowHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShow(ctx, req.ProjectId, req.Id, req.Params, authorizationHeader))
}

type WikiShowPdfRequest struct {
	ProjectId string                    `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                    `json:"id" jsonschema:"description=The title of the wiki."`
	Params    *client.WikiShowPdfParams `json:"params,omitempty"`
}

func registerWikiShowPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_pdf",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowPdfHandler))
}

func wikiShowPdfHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowPdf(ctx, req.ProjectId, req.Id, req.Params, authorizationHeader))
}

type WikiShowTxtRequest struct {
	ProjectId string                    `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                    `json:"id" jsonschema:"description=The title of the wiki."`
	Params    *client.WikiShowTxtParams `json:"params,omitempty"`
}

func registerWikiShowTxt(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowTxtRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_txt",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowTxtHandler))
}

func wikiShowTxtHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowTxtRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowTxt(ctx, req.ProjectId, req.Id, req.Params, authorizationHeader))
}

type WikiShowVersionRequest struct {
	ProjectId string                        `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                        `json:"id" jsonschema:"description=The title of the wiki."`
	Version   int                           `json:"version" jsonschema:"description=The version of the wiki."`
	Params    *client.WikiShowVersionParams `json:"params,omitempty"`
}

func registerWikiShowVersion(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowVersionRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_version",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowVersionHandler))
}

func wikiShowVersionHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowVersionRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowVersion(ctx, req.ProjectId, req.Id, req.Version, req.Params, authorizationHeader))
}

type WikiShowVersionPdfRequest struct {
	ProjectId string                           `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                           `json:"id" jsonschema:"description=The title of the wiki."`
	Version   int                              `json:"version" jsonschema:"description=The version of the wiki."`
	Params    *client.WikiShowVersionPdfParams `json:"params,omitempty"`
}

func registerWikiShowVersionPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowVersionPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_version_pdf",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowVersionPdfHandler))
}

func wikiShowVersionPdfHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowVersionPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowVersionPdf(ctx, req.ProjectId, req.Id, req.Version, req.Params, authorizationHeader))
}

type WikiShowVersionTxtRequest struct {
	ProjectId string                           `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Id        string                           `json:"id" jsonschema:"description=The title of the wiki."`
	Version   int                              `json:"version" jsonschema:"description=The version of the wiki."`
	Params    *client.WikiShowVersionTxtParams `json:"params,omitempty"`
}

func registerWikiShowVersionTxt(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&WikiShowVersionTxtRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("wiki_show_version_txt",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(wikiShowVersionTxtHandler))
}

func wikiShowVersionTxtHandler(ctx context.Context, request mcp.CallToolRequest, req WikiShowVersionTxtRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.WikiShowVersionTxt(ctx, req.ProjectId, req.Id, req.Version, req.Params, authorizationHeader))
}
