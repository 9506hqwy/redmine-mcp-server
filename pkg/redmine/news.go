package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerNewsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("news_index",
		mcp.WithDescription("Returns all news items across all projects with pagination."),
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

	s.AddTool(tool, newsIndexHandler)
}

func newsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseNewsIndex(request)
	return toResult(c.NewsIndex(ctx, &params, authorizationHeader))
}

func parseNewsIndex(request mcp.CallToolRequest) client.NewsIndexParams {
	params := client.NewsIndexParams{}

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

func registerNewsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("news_destroy",
		mcp.WithDescription("Deletes the news with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the news."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, newsDestroyHandler)
}

func newsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseNewsDestroy(request)
	return toResult(c.NewsDestroy(ctx, id, &params, authorizationHeader))
}

func parseNewsDestroy(request mcp.CallToolRequest) client.NewsDestroyParams {
	params := client.NewsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerNewsShow(s *server.MCPServer) {
	tool := mcp.NewTool("news_show",
		mcp.WithDescription("Returns the news item with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the news."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`, `comments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, newsShowHandler)
}

func newsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseNewsShow(request)
	return toResult(c.NewsShow(ctx, id, &params, authorizationHeader))
}

func parseNewsShow(request mcp.CallToolRequest) client.NewsShowParams {
	params := client.NewsShowParams{}

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

func registerNewsIndexProject(s *server.MCPServer) {
	tool := mcp.NewTool("news_index_project",
		mcp.WithDescription("Returns all news items across all projects with pagination."),
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
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, newsIndexProjectHandler)
}

func newsIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseNewsIndexProject(request)
	return toResult(c.NewsIndexProject(ctx, project_id, &params, authorizationHeader))
}

func parseNewsIndexProject(request mcp.CallToolRequest) client.NewsIndexProjectParams {
	params := client.NewsIndexProjectParams{}

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
