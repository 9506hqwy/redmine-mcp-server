package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerIssueCategoriesDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("issue_categories_destroy",
		mcp.WithDescription("Deletes the issue category with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue category."),
			mcp.Required(),
		),
		mcp.WithNumber("reassign_to_id",
			mcp.Description("The ID of the issue category."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueCategoriesDestroyHandler)
}

func issueCategoriesDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssueCategoriesDestroy(request)
	return toResult(c.IssueCategoriesDestroy(ctx, id, &params, authorizationHeader))
}

func parseIssueCategoriesDestroy(request mcp.CallToolRequest) client.IssueCategoriesDestroyParams {
	params := client.IssueCategoriesDestroyParams{}

	reassign_to_id := request.GetInt("reassign_to_id", math.MinInt)
	if reassign_to_id != math.MinInt {

		params.ReassignToId = &reassign_to_id
	}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerIssueCategoriesShow(s *server.MCPServer) {
	tool := mcp.NewTool("issue_categories_show",
		mcp.WithDescription("Returns the issue category with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue category."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueCategoriesShowHandler)
}

func issueCategoriesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssueCategoriesShow(request)
	return toResult(c.IssueCategoriesShow(ctx, id, &params, authorizationHeader))
}

func parseIssueCategoriesShow(request mcp.CallToolRequest) client.IssueCategoriesShowParams {
	params := client.IssueCategoriesShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerIssueCategoriesIndex(s *server.MCPServer) {
	tool := mcp.NewTool("issue_categories_index",
		mcp.WithDescription("Returns the issue categories available for the specified project by ID or identifier."),
		mcp.WithNumber("X-Redmine-Nometa",
			mcp.Description("If set to 1, the response will not include metadata information."),

			mcp.Enum("1"),
		),
		mcp.WithNumber("nometa",
			mcp.Description("If set to 1, the response will not include metadata information."),

			mcp.Enum("1"),
		),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issueCategoriesIndexHandler)
}

func issueCategoriesIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseIssueCategoriesIndex(request)
	return toResult(c.IssueCategoriesIndex(ctx, project_id, &params, authorizationHeader))
}

func parseIssueCategoriesIndex(request mcp.CallToolRequest) client.IssueCategoriesIndexParams {
	params := client.IssueCategoriesIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	nometa := request.GetInt("nometa", math.MinInt)
	if nometa != math.MinInt {

		params.Nometa = &nometa
	}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
