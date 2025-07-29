package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerRepositoriesRemoveRelatedIssue(s *server.MCPServer) {
	tool := mcp.NewTool("repositories_remove_related_issue",
		mcp.WithDescription("Remove a related issue from the specified revision."),
		mcp.WithString("id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("repository_id",
			mcp.Description("The unique identifier of the repository."),
			mcp.Required(),
		),
		mcp.WithString("rev",
			mcp.Description("The revision identifier of the changeset."),
			mcp.Required(),
		),
		mcp.WithNumber("issue_id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, repositoriesRemoveRelatedIssueHandler)
}

func repositoriesRemoveRelatedIssueHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	repository_id := request.GetString("repository_id", "")
	rev := request.GetString("rev", "")
	issue_id := request.GetInt("issue_id", math.MinInt)
	params := parseRepositoriesRemoveRelatedIssue(request)
	return toResult(c.RepositoriesRemoveRelatedIssue(ctx, id, repository_id, rev, issue_id, &params, authorizationHeader))
}

func parseRepositoriesRemoveRelatedIssue(request mcp.CallToolRequest) client.RepositoriesRemoveRelatedIssueParams {
	params := client.RepositoriesRemoveRelatedIssueParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
