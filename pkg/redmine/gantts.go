package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type GanttsShowPdfRequest struct {
	Params *client.GanttsShowPdfParams `json:"params,omitempty"`
}

func registerGanttsShowPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GanttsShowPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("gantts_show_pdf",
		mcp.WithDescription("Download the Gantt chart."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ganttsShowPdfHandler))
}

func ganttsShowPdfHandler(ctx context.Context, request mcp.CallToolRequest, req GanttsShowPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GanttsShowPdf(ctx, req.Params, authorizationHeader))
}

type GanttsShowPngRequest struct {
	Params *client.GanttsShowPngParams `json:"params,omitempty"`
}

func registerGanttsShowPng(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GanttsShowPngRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("gantts_show_png",
		mcp.WithDescription("Download the Gantt chart."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ganttsShowPngHandler))
}

func ganttsShowPngHandler(ctx context.Context, request mcp.CallToolRequest, req GanttsShowPngRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GanttsShowPng(ctx, req.Params, authorizationHeader))
}

type GanttsShowProjectPdfRequest struct {
	ProjectId string                             `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.GanttsShowProjectPdfParams `json:"params,omitempty"`
}

func registerGanttsShowProjectPdf(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GanttsShowProjectPdfRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("gantts_show_project_pdf",
		mcp.WithDescription("Download the Gantt chart for the specified project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ganttsShowProjectPdfHandler))
}

func ganttsShowProjectPdfHandler(ctx context.Context, request mcp.CallToolRequest, req GanttsShowProjectPdfRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GanttsShowProjectPdf(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type GanttsShowProjectPngRequest struct {
	ProjectId string                             `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.GanttsShowProjectPngParams `json:"params,omitempty"`
}

func registerGanttsShowProjectPng(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&GanttsShowProjectPngRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("gantts_show_project_png",
		mcp.WithDescription("Download the Gantt chart for the specified project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(ganttsShowProjectPngHandler))
}

func ganttsShowProjectPngHandler(ctx context.Context, request mcp.CallToolRequest, req GanttsShowProjectPngRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.GanttsShowProjectPng(ctx, req.ProjectId, req.Params, authorizationHeader))
}
