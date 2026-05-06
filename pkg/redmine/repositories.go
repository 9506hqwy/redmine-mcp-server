package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type RepositoriesAddRelatedIssueRequest struct {
	Id           string                                            `json:"id" jsonschema:"description=The ID or identifier of the project."`
	RepositoryId string                                            `json:"repository_id" jsonschema:"description=The unique identifier of the repository."`
	Rev          string                                            `json:"rev" jsonschema:"description=The revision identifier of the changeset."`
	Params       *client.RepositoriesAddRelatedIssueParams         `json:"params,omitempty"`
	Body         client.RepositoriesAddRelatedIssueJSONRequestBody `json:"body"`
}

func registerRepositoriesAddRelatedIssue(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&RepositoriesAddRelatedIssueRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("repositories_add_related_issue",
		mcp.WithDescription("Add a related issue to the specified revision."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(repositoriesAddRelatedIssueHandler))
}

func repositoriesAddRelatedIssueHandler(ctx context.Context, request mcp.CallToolRequest, req RepositoriesAddRelatedIssueRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.RepositoriesAddRelatedIssue(ctx, req.Id, req.RepositoryId, req.Rev, req.Params, req.Body, authorizationHeader))
}

type RepositoriesRemoveRelatedIssueRequest struct {
	Id           string                                       `json:"id" jsonschema:"description=The ID or identifier of the project."`
	RepositoryId string                                       `json:"repository_id" jsonschema:"description=The unique identifier of the repository."`
	Rev          string                                       `json:"rev" jsonschema:"description=The revision identifier of the changeset."`
	IssueId      int                                          `json:"issue_id" jsonschema:"description=The ID of the issue."`
	Params       *client.RepositoriesRemoveRelatedIssueParams `json:"params,omitempty"`
}

func registerRepositoriesRemoveRelatedIssue(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&RepositoriesRemoveRelatedIssueRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("repositories_remove_related_issue",
		mcp.WithDescription("Remove a related issue from the specified revision."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(repositoriesRemoveRelatedIssueHandler))
}

func repositoriesRemoveRelatedIssueHandler(ctx context.Context, request mcp.CallToolRequest, req RepositoriesRemoveRelatedIssueRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.RepositoriesRemoveRelatedIssue(ctx, req.Id, req.RepositoryId, req.Rev, req.IssueId, req.Params, authorizationHeader))
}
