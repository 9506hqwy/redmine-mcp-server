package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerWatchersDestroyIssue(s *server.MCPServer) {
	tool := mcp.NewTool("watchers_destroy_issue",
		mcp.WithDescription("Deletes the watcher with the specified ID from the issue."),
		mcp.WithNumber("issue_id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The ID of the user."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, watchersDestroyIssueHandler)
}

func watchersDestroyIssueHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	issue_id := request.GetInt("issue_id", math.MinInt)
	user_id := request.GetInt("user_id", math.MinInt)
	params := parseWatchersDestroyIssue(request)
	return toResult(c.WatchersDestroyIssue(ctx, issue_id, user_id, &params, authorizationHeader))
}

func parseWatchersDestroyIssue(request mcp.CallToolRequest) client.WatchersDestroyIssueParams {
	params := client.WatchersDestroyIssueParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerWatchersDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("watchers_destroy",
		mcp.WithDescription("Deletes the watcher with the specified ID."),
		mcp.WithString("object_type",
			mcp.Description("The type of the object to be watched."),
			mcp.Required(),
		),
		mcp.WithNumber("object_id",
			mcp.Description("The ID of the object."),
			mcp.Required(),
		),
		mcp.WithNumber("user_id",
			mcp.Description("The ID of the user."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, watchersDestroyHandler)
}

func watchersDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseWatchersDestroy(request)
	return toResult(c.WatchersDestroy(ctx, &params, authorizationHeader))
}

func parseWatchersDestroy(request mcp.CallToolRequest) client.WatchersDestroyParams {
	params := client.WatchersDestroyParams{}

	object_type := request.GetString("object_type", "")
	if object_type != "" {

		params.ObjectType = object_type
	}

	object_id := request.GetInt("object_id", math.MinInt)
	if object_id != math.MinInt {

		params.ObjectId = object_id
	}

	user_id := request.GetInt("user_id", math.MinInt)
	if user_id != math.MinInt {

		params.UserId = user_id
	}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
