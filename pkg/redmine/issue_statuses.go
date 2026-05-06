package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type IssueStatusesIndexRequest struct {
	Params *client.IssueStatusesIndexParams `json:"params,omitempty"`
}

func registerIssueStatusesIndex(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&IssueStatusesIndexRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("issue_statuses_index",
		mcp.WithDescription("Returns a list of all issue statuses."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(issueStatusesIndexHandler))
}

func issueStatusesIndexHandler(ctx context.Context, request mcp.CallToolRequest, req IssueStatusesIndexRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.IssueStatusesIndex(ctx, req.Params, authorizationHeader))
}
