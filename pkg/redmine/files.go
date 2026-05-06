package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type FilesIndexRequest struct {
	ProjectId string                   `json:"project_id" jsonschema:"description=The ID or identifier of the project."`
	Params    *client.FilesIndexParams `json:"params,omitempty"`
}

func registerFilesIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&FilesIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("files_index",
		mcp.WithDescription("Returns a list of all files."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(filesIndexHandler))
}

func filesIndexHandler(ctx context.Context, request mcp.CallToolRequest, req FilesIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.FilesIndex(ctx, req.ProjectId, req.Params, authorizationHeader))
}
