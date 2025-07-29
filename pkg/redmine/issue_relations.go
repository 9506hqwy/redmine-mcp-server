package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerIssueRelationsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("issue_relations_index",
		mcp.WithDescription("Returns the relations for the specified issue ID."),
		mcp.WithNumber("issue_id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueRelationsIndexHandler)
}

func issueRelationsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	issue_id := request.GetInt("issue_id", math.MinInt)
	params := parseIssueRelationsIndex(request)
	return toResult(c.IssueRelationsIndex(ctx, issue_id, &params, authorizationHeader))
}

func parseIssueRelationsIndex(request mcp.CallToolRequest) client.IssueRelationsIndexParams {
	params := client.IssueRelationsIndexParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerIssueRelationsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("issue_relations_destroy",
		mcp.WithDescription("Deletes the relation with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the relation."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueRelationsDestroyHandler)
}

func issueRelationsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssueRelationsDestroy(request)
	return toResult(c.IssueRelationsDestroy(ctx, id, &params, authorizationHeader))
}

func parseIssueRelationsDestroy(request mcp.CallToolRequest) client.IssueRelationsDestroyParams {
	params := client.IssueRelationsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerIssueRelationsShow(s *server.MCPServer) {
	tool := mcp.NewTool("issue_relations_show",
		mcp.WithDescription("Returns the relation with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the relation."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueRelationsShowHandler)
}

func issueRelationsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssueRelationsShow(request)
	return toResult(c.IssueRelationsShow(ctx, id, &params, authorizationHeader))
}

func parseIssueRelationsShow(request mcp.CallToolRequest) client.IssueRelationsShowParams {
	params := client.IssueRelationsShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
