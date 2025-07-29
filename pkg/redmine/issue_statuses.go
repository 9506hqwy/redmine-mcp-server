package redmine

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerIssueStatusesIndex(s *server.MCPServer) {
	tool := mcp.NewTool("issue_statuses_index",
		mcp.WithDescription("Returns a list of all issue statuses."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueStatusesIndexHandler)
}

func issueStatusesIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssueStatusesIndex(request)
	return toResult(c.IssueStatusesIndex(ctx, &params, authorizationHeader))
}

func parseIssueStatusesIndex(request mcp.CallToolRequest) client.IssueStatusesIndexParams {
	params := client.IssueStatusesIndexParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
