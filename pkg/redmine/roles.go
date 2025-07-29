package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerRolesIndex(s *server.MCPServer) {
	tool := mcp.NewTool("roles_index",
		mcp.WithDescription("Returns a list of all roles."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, rolesIndexHandler)
}

func rolesIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseRolesIndex(request)
	return toResult(c.RolesIndex(ctx, &params, authorizationHeader))
}

func parseRolesIndex(request mcp.CallToolRequest) client.RolesIndexParams {
	params := client.RolesIndexParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerRolesShow(s *server.MCPServer) {
	tool := mcp.NewTool("roles_show",
		mcp.WithDescription("Returns the role with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the role."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, rolesShowHandler)
}

func rolesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseRolesShow(request)
	return toResult(c.RolesShow(ctx, id, &params, authorizationHeader))
}

func parseRolesShow(request mcp.CallToolRequest) client.RolesShowParams {
	params := client.RolesShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
