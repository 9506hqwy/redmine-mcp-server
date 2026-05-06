package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type JournalsUpdatePatchRequest struct {
	Id     int                                       `json:"id" jsonschema:"description=The ID of the journal."`
	Params *client.JournalsUpdatePatchParams         `json:"params,omitempty"`
	Body   client.JournalsUpdatePatchJSONRequestBody `json:"body,omitempty"`
}

func registerJournalsUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&JournalsUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("journals_update_patch",
		mcp.WithDescription("Update the journal with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(journalsUpdatePatchHandler))
}

func journalsUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req JournalsUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.JournalsUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}
