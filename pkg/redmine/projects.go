package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerProjectsIndexCsv(s *server.MCPServer) {
	tool := mcp.NewTool("projects_index_csv",
		mcp.WithDescription("Returns all projects (including all public projects and private projects to which the user has access)."),
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
		mcp.WithString("query.status",
			mcp.Description("The expression of status of the project. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `1`: active - `5`: closed examples: - active or closed: `1|5` - not closed: `!5`"),
		),
		mcp.WithString("query.id",
			mcp.Description("The expression of identifier of the project. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.name",
			mcp.Description("The expression of name of the project. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the project. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent project id of the project. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.is_public",
			mcp.Description("The expression of visibility of the project. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `0`: private - `1`: public examples: - public only: `1` - public or private: `0|1`"),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the project. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the project. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `trackers`, `issue_categories`, `enabled_modules`, `time_entry_activities`, `issue_custom_fields`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsIndexCsvHandler)
}

func projectsIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsIndexCsv(request)
	return toResult(c.ProjectsIndexCsv(ctx, &params, authorizationHeader))
}

func parseProjectsIndexCsv(request mcp.CallToolRequest) client.ProjectsIndexCsvParams {
	params := client.ProjectsIndexCsvParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseProjectsIndexCsvQuery(&request)

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

func parseProjectsIndexCsvQuery(request *mcp.CallToolRequest) *client.ProjectsIndexCsvParams_Query {
	params := client.ProjectsIndexCsvParams_Query{}

	status := request.GetString("query.status", "")
	if status != "" {

		params.Status = &status
	}

	id := request.GetString("query.id", "")
	if id != "" {

		params.Id = &id
	}

	name := request.GetString("query.name", "")
	if name != "" {

		params.Name = &name
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	is_public := request.GetString("query.is_public", "")
	if is_public != "" {

		params.IsPublic = &is_public
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	return &params
}

func registerProjectsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("projects_index",
		mcp.WithDescription("Returns all projects (including all public projects and private projects to which the user has access)."),
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
		mcp.WithString("query.status",
			mcp.Description("The expression of status of the project. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `1`: active - `5`: closed examples: - active or closed: `1|5` - not closed: `!5`"),
		),
		mcp.WithString("query.id",
			mcp.Description("The expression of identifier of the project. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.name",
			mcp.Description("The expression of name of the project. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the project. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent project id of the project. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.is_public",
			mcp.Description("The expression of visibility of the project. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `0`: private - `1`: public examples: - public only: `1` - public or private: `0|1`"),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the project. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the project. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `trackers`, `issue_categories`, `enabled_modules`, `time_entry_activities`, `issue_custom_fields`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsIndexHandler)
}

func projectsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseProjectsIndex(request)
	return toResult(c.ProjectsIndex(ctx, &params, authorizationHeader))
}

func parseProjectsIndex(request mcp.CallToolRequest) client.ProjectsIndexParams {
	params := client.ProjectsIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseProjectsIndexQuery(&request)

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

func parseProjectsIndexQuery(request *mcp.CallToolRequest) *client.ProjectsIndexParams_Query {
	params := client.ProjectsIndexParams_Query{}

	status := request.GetString("query.status", "")
	if status != "" {

		params.Status = &status
	}

	id := request.GetString("query.id", "")
	if id != "" {

		params.Id = &id
	}

	name := request.GetString("query.name", "")
	if name != "" {

		params.Name = &name
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	is_public := request.GetString("query.is_public", "")
	if is_public != "" {

		params.IsPublic = &is_public
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	return &params
}

func registerProjectsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("projects_destroy",
		mcp.WithDescription("Deletes the project with the specified ID or identifier."),
		mcp.WithString("id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsDestroyHandler)
}

func projectsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseProjectsDestroy(request)
	return toResult(c.ProjectsDestroy(ctx, id, &params, authorizationHeader))
}

func parseProjectsDestroy(request mcp.CallToolRequest) client.ProjectsDestroyParams {
	params := client.ProjectsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerProjectsShow(s *server.MCPServer) {
	tool := mcp.NewTool("projects_show",
		mcp.WithDescription("Returns the project with the specified ID or identifier."),
		mcp.WithString("id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `trackers`, `issue_categories`, `enabled_modules`, `time_entry_activities`, `issue_custom_fields`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsShowHandler)
}

func projectsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseProjectsShow(request)
	return toResult(c.ProjectsShow(ctx, id, &params, authorizationHeader))
}

func parseProjectsShow(request mcp.CallToolRequest) client.ProjectsShowParams {
	params := client.ProjectsShowParams{}

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

func registerProjectsArchivePost(s *server.MCPServer) {
	tool := mcp.NewTool("projects_archive_post",
		mcp.WithDescription("Archives the project with the specified ID or identifier."),
		mcp.WithString("id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsArchivePostHandler)
}

func projectsArchivePostHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseProjectsArchivePost(request)
	return toResult(c.ProjectsArchivePost(ctx, id, &params, authorizationHeader))
}

func parseProjectsArchivePost(request mcp.CallToolRequest) client.ProjectsArchivePostParams {
	params := client.ProjectsArchivePostParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerProjectsUnarchivePost(s *server.MCPServer) {
	tool := mcp.NewTool("projects_unarchive_post",
		mcp.WithDescription("Unarchives the project with the specified ID or identifier."),
		mcp.WithString("id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, projectsUnarchivePostHandler)
}

func projectsUnarchivePostHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseProjectsUnarchivePost(request)
	return toResult(c.ProjectsUnarchivePost(ctx, id, &params, authorizationHeader))
}

func parseProjectsUnarchivePost(request mcp.CallToolRequest) client.ProjectsUnarchivePostParams {
	params := client.ProjectsUnarchivePostParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
