package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerGroupsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("groups_index",
		mcp.WithDescription("Returns a list of all groups."),
		mcp.WithNumber("X-Redmine-Nometa",
			mcp.Description("If set to 1, the response will not include metadata information."),

			mcp.Enum("1"),
		),
		mcp.WithNumber("pagination.offset",
			mcp.Description("The offset of the first object to retrieve If not specified, it defaults to 0. (default: 0)"),
		),
		mcp.WithNumber("pagination.limit",
			mcp.Description("The number of items to be present in the response. If not specified, it defaults to 25. (default: 25)"),
		),
		mcp.WithNumber("pagination.nometa",
			mcp.Description("If set to 1, the response will not include pagination information."),

			mcp.Enum("1"),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, groupsIndexHandler)
}

func groupsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGroupsIndex(request)
	return toResult(c.GroupsIndex(ctx, &params, authorizationHeader))
}

func parseGroupsIndex(request mcp.CallToolRequest) client.GroupsIndexParams {
	params := client.GroupsIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGroupsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("groups_destroy",
		mcp.WithDescription("Deletes the group with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the group."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, groupsDestroyHandler)
}

func groupsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseGroupsDestroy(request)
	return toResult(c.GroupsDestroy(ctx, id, &params, authorizationHeader))
}

func parseGroupsDestroy(request mcp.CallToolRequest) client.GroupsDestroyParams {
	params := client.GroupsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGroupsShow(s *server.MCPServer) {
	tool := mcp.NewTool("groups_show",
		mcp.WithDescription("Returns the group with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the group."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `users`, `memberships`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, groupsShowHandler)
}

func groupsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseGroupsShow(request)
	return toResult(c.GroupsShow(ctx, id, &params, authorizationHeader))
}

func parseGroupsShow(request mcp.CallToolRequest) client.GroupsShowParams {
	params := client.GroupsShowParams{}

	include := request.GetString("include", "")
	if include != "" {
		include := strings.Split(include, ",")
		params.Include = &include
	}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGroupsRemoveUser(s *server.MCPServer) {
	tool := mcp.NewTool("groups_remove_user",
		mcp.WithDescription("Removes a user from a group."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the group."),
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

	s.AddTool(tool, groupsRemoveUserHandler)
}

func groupsRemoveUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	user_id := request.GetInt("user_id", math.MinInt)
	params := parseGroupsRemoveUser(request)
	return toResult(c.GroupsRemoveUser(ctx, id, user_id, &params, authorizationHeader))
}

func parseGroupsRemoveUser(request mcp.CallToolRequest) client.GroupsRemoveUserParams {
	params := client.GroupsRemoveUserParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
