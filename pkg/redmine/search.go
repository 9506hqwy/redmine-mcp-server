package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerSearchIndexProject(s *server.MCPServer) {
	tool := mcp.NewTool("search_index_project",
		mcp.WithDescription("Returns search results based on the specified query parameters."),
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
		mcp.WithString("q",
			mcp.Description("Query string. Multiple values can be specified, separated by a space (\" \")."),
			mcp.Required(),
		),
		mcp.WithString("query.scope",
			mcp.Description("Search scope condition. Possible values are: - `all`: Search all projects - `my_project`: Search assigned projects - `bookmarks`: Search bookmarked projects - `subprojects`: Include subproject when project specified"),

			mcp.Enum("all", "my_project", "bookmarks", "subprojects"),
		),
		mcp.WithBoolean("query.all_words",
			mcp.Description("matched all query strings or not."),
		),
		mcp.WithBoolean("query.titles_only",
			mcp.Description("matched only title or not."),
		),
		mcp.WithBoolean("query.issues",
			mcp.Description("Include issues or not."),
		),
		mcp.WithBoolean("query.news",
			mcp.Description("Include news or not."),
		),
		mcp.WithBoolean("query.documents",
			mcp.Description("Include documents or not."),
		),
		mcp.WithBoolean("query.changesets",
			mcp.Description("Include changesets or not."),
		),
		mcp.WithBoolean("query.wiki_pages",
			mcp.Description("Include documents or not."),
		),
		mcp.WithBoolean("query.messages",
			mcp.Description("Include messages or not."),
		),
		mcp.WithBoolean("query.projects",
			mcp.Description("Include projects or not."),
		),
		mcp.WithBoolean("query.open_issues",
			mcp.Description("Filterd by open issues."),
		),
		mcp.WithString("query.attachments",
			mcp.Description("Filterd by description and attachment. - `0`: Seach only in description - `1`: Search by description and attachment - `only`: Search only in attachment"),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, searchIndexProjectHandler)
}

func searchIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseSearchIndexProject(request)
	return toResult(c.SearchIndexProject(ctx, project_id, &params, authorizationHeader))
}

func parseSearchIndexProject(request mcp.CallToolRequest) client.SearchIndexProjectParams {
	params := client.SearchIndexProjectParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	q := request.GetString("q", "")
	if q != "" {

		params.Q = q
	}

	params.Query = parseSearchIndexProjectQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerSearchIndex(s *server.MCPServer) {
	tool := mcp.NewTool("search_index",
		mcp.WithDescription("Returns search results based on the specified query parameters."),
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
		mcp.WithString("q",
			mcp.Description("Query string. Multiple values can be specified, separated by a space (\" \")."),
			mcp.Required(),
		),
		mcp.WithString("query.scope",
			mcp.Description("Search scope condition. Possible values are: - `all`: Search all projects - `my_project`: Search assigned projects - `bookmarks`: Search bookmarked projects - `subprojects`: Include subproject when project specified"),

			mcp.Enum("all", "my_project", "bookmarks", "subprojects"),
		),
		mcp.WithBoolean("query.all_words",
			mcp.Description("matched all query strings or not."),
		),
		mcp.WithBoolean("query.titles_only",
			mcp.Description("matched only title or not."),
		),
		mcp.WithBoolean("query.issues",
			mcp.Description("Include issues or not."),
		),
		mcp.WithBoolean("query.news",
			mcp.Description("Include news or not."),
		),
		mcp.WithBoolean("query.documents",
			mcp.Description("Include documents or not."),
		),
		mcp.WithBoolean("query.changesets",
			mcp.Description("Include changesets or not."),
		),
		mcp.WithBoolean("query.wiki_pages",
			mcp.Description("Include documents or not."),
		),
		mcp.WithBoolean("query.messages",
			mcp.Description("Include messages or not."),
		),
		mcp.WithBoolean("query.projects",
			mcp.Description("Include projects or not."),
		),
		mcp.WithBoolean("query.open_issues",
			mcp.Description("Filterd by open issues."),
		),
		mcp.WithString("query.attachments",
			mcp.Description("Filterd by description and attachment. - `0`: Seach only in description - `1`: Search by description and attachment - `only`: Search only in attachment"),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, searchIndexHandler)
}

func searchIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseSearchIndex(request)
	return toResult(c.SearchIndex(ctx, &params, authorizationHeader))
}

func parseSearchIndex(request mcp.CallToolRequest) client.SearchIndexParams {
	params := client.SearchIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	q := request.GetString("q", "")
	if q != "" {

		params.Q = q
	}

	params.Query = parseSearchIndexQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
