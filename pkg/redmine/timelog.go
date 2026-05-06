package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type TimelogIndexProjectCsvRequest struct {
	ProjectId string                               `json:"project_id" jsonschema:"description=The ID or identifer of the project."`
	Params    *client.TimelogIndexProjectCsvParams `json:"params,omitempty"`
}

func registerTimelogIndexProjectCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogIndexProjectCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_index_project_csv",
		mcp.WithDescription("Returns a list of time entries for the specified project in CSV format."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogIndexProjectCsvHandler))
}

func timelogIndexProjectCsvHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogIndexProjectCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogIndexProjectCsv(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type TimelogIndexProjectRequest struct {
	ProjectId string                            `json:"project_id" jsonschema:"description=The ID or identifer of the project."`
	Params    *client.TimelogIndexProjectParams `json:"params,omitempty"`
}

func registerTimelogIndexProject(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogIndexProjectRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_index_project",
		mcp.WithDescription("Returns a list of time entries for the specified project."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogIndexProjectHandler))
}

func timelogIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogIndexProjectRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogIndexProject(ctx, req.ProjectId, req.Params, authorizationHeader))
}

type TimelogIndexCsvRequest struct {
	Params *client.TimelogIndexCsvParams `json:"params,omitempty"`
}

func registerTimelogIndexCsv(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogIndexCsvRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_index_csv",
		mcp.WithDescription("Returns a list of time entries."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogIndexCsvHandler))
}

func timelogIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogIndexCsvRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogIndexCsv(ctx, req.Params, authorizationHeader))
}

type TimelogIndexRequest struct {
	Params *client.TimelogIndexParams `json:"params,omitempty"`
}

func registerTimelogIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_index",
		mcp.WithDescription("Returns a list of time entries."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogIndexHandler))
}

func timelogIndexHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogIndex(ctx, req.Params, authorizationHeader))
}

type TimelogDestroyRequest struct {
	Id     int                          `json:"id" jsonschema:"description=The ID of the time entry."`
	Params *client.TimelogDestroyParams `json:"params,omitempty"`
}

func registerTimelogDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_destroy",
		mcp.WithDescription("Deletes the time entry with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogDestroyHandler))
}

func timelogDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type TimelogShowRequest struct {
	Id     int                       `json:"id" jsonschema:"description=The ID of the time entry."`
	Params *client.TimelogShowParams `json:"params,omitempty"`
}

func registerTimelogShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&TimelogShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("timelog_show",
		mcp.WithDescription("Returns the time entry with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(timelogShowHandler))
}

func timelogShowHandler(ctx context.Context, request mcp.CallToolRequest, req TimelogShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.TimelogShow(ctx, req.Id, req.Params, authorizationHeader))
}
