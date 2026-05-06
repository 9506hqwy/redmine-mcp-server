package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type VersionsCreateRequest struct {
	ProjectId string                               `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.VersionsCreateParams         `json:"params,omitempty"`
	Body      client.VersionsCreateJSONRequestBody `json:"body"`
}

func registerVersionsCreate(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsCreateRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_create",
		mcp.WithDescription("Creates a new version for the project with the specified ID or identifier (:project_id)."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsCreateHandler))
}

func versionsCreateHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsCreateRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsCreate(ctx, req.ProjectId, req.Params, req.Body, authorizationHeader))
}

type VersionsIndexRequest struct {
	ProjectId string                      `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.VersionsIndexParams `json:"params,omitempty"`
}

func registerVersionsIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_index",
		mcp.WithDescription("Returns the versions available for the project with the specified ID or identifier (:project_id). The response may include shared versions from other projects."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsIndexHandler))
}

func versionsIndexHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsIndex(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type VersionsShowTxtRequest struct {
	Id     int                           `json:"id" jsonschema:"description=The ID of the version."`
	Params *client.VersionsShowTxtParams `json:"params,omitempty"`
}

func registerVersionsShowTxt(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsShowTxtRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_show_txt",
		mcp.WithDescription("Returns the version with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsShowTxtHandler))
}

func versionsShowTxtHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsShowTxtRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsShowTxt(ctx, req.Id, req.Params, authorizationHeader))
}

type VersionsDestroyRequest struct {
	Id     int                           `json:"id" jsonschema:"description=The ID of the version."`
	Params *client.VersionsDestroyParams `json:"params,omitempty"`
}

func registerVersionsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_destroy",
		mcp.WithDescription("Deletes the version with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsDestroyHandler))
}

func versionsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type VersionsUpdatePatchRequest struct {
	Id     int                                       `json:"id" jsonschema:"description=The ID of the version."`
	Params *client.VersionsUpdatePatchParams         `json:"params,omitempty"`
	Body   client.VersionsUpdatePatchJSONRequestBody `json:"body"`
}

func registerVersionsUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_update_patch",
		mcp.WithDescription("Updates the version with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsUpdatePatchHandler))
}

func versionsUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}

type VersionsShowRequest struct {
	Id     int                        `json:"id" jsonschema:"description=The ID of the version."`
	Params *client.VersionsShowParams `json:"params,omitempty"`
}

func registerVersionsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&VersionsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("versions_show",
		mcp.WithDescription("Returns the version with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(versionsShowHandler))
}

func versionsShowHandler(ctx context.Context, request mcp.CallToolRequest, req VersionsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.VersionsShow(ctx, req.Id, req.Params, authorizationHeader))
}
