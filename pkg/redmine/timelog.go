package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerTimelogIndexProjectCsv(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_index_project_csv",
		mcp.WithDescription("Returns a list of time entries for the specified project in CSV format."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifer of the project."),
			mcp.Required(),
		),
		mcp.WithString("query.spent_on",
			mcp.Description("The expression of spent_on of the time entry. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.tracker_id",
			mcp.Description("The expression of issue tracker id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.parent_id",
			mcp.Description("The expression of issue parent id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.status_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.fixed_version_id",
			mcp.Description("The expression of issue fixed version id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.category_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.subject",
			mcp.Description("The expression of issue subject of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.user_id",
			mcp.Description("The expression of user id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.group",
			mcp.Description("The expression of user group of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.role",
			mcp.Description("The expression of user role of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.activity_id",
			mcp.Description("The expression of activity id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.comments",
			mcp.Description("The expression of comments of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.hours",
			mcp.Description("The expression of hours of the time entry. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogIndexProjectCsvHandler)
}

func timelogIndexProjectCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseTimelogIndexProjectCsv(request)
	return toResult(c.TimelogIndexProjectCsv(ctx, project_id, &params, authorizationHeader))
}

func parseTimelogIndexProjectCsv(request mcp.CallToolRequest) client.TimelogIndexProjectCsvParams {
	params := client.TimelogIndexProjectCsvParams{}

	params.Query = parseTimelogIndexProjectCsvQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func parseTimelogIndexProjectCsvQuery(request *mcp.CallToolRequest) *client.TimelogIndexProjectCsvParams_Query {
	params := client.TimelogIndexProjectCsvParams_Query{}

	spent_on := request.GetString("query.spent_on", "")
	if spent_on != "" {

		params.SpentOn = &spent_on
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	issue_tracker_id := request.GetString("query.issue.tracker_id", "")
	if issue_tracker_id != "" {

		params.IssueTrackerId = &issue_tracker_id
	}

	issue_parent_id := request.GetString("query.issue.parent_id", "")
	if issue_parent_id != "" {

		params.IssueParentId = &issue_parent_id
	}

	issue_status_id := request.GetString("query.issue.status_id", "")
	if issue_status_id != "" {

		params.IssueStatusId = &issue_status_id
	}

	issue_fixed_version_id := request.GetString("query.issue.fixed_version_id", "")
	if issue_fixed_version_id != "" {

		params.IssueFixedVersionId = &issue_fixed_version_id
	}

	issue_category_id := request.GetString("query.issue.category_id", "")
	if issue_category_id != "" {

		params.IssueCategoryId = &issue_category_id
	}

	issue_subject := request.GetString("query.issue.subject", "")
	if issue_subject != "" {

		params.IssueSubject = &issue_subject
	}

	user_id := request.GetString("query.user_id", "")
	if user_id != "" {

		params.UserId = &user_id
	}

	user_group := request.GetString("query.user.group", "")
	if user_group != "" {

		params.UserGroup = &user_group
	}

	user_role := request.GetString("query.user.role", "")
	if user_role != "" {

		params.UserRole = &user_role
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	activity_id := request.GetString("query.activity_id", "")
	if activity_id != "" {

		params.ActivityId = &activity_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	issue_comments := request.GetString("query.issue.comments", "")
	if issue_comments != "" {

		params.IssueComments = &issue_comments
	}

	hours := request.GetString("query.hours", "")
	if hours != "" {

		params.Hours = &hours
	}

	return &params
}

func registerTimelogIndexProject(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_index_project",
		mcp.WithDescription("Returns a list of time entries for the specified project."),
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
			mcp.Description("The ID or identifer of the project."),
			mcp.Required(),
		),
		mcp.WithString("query.spent_on",
			mcp.Description("The expression of spent_on of the time entry. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.tracker_id",
			mcp.Description("The expression of issue tracker id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.parent_id",
			mcp.Description("The expression of issue parent id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.status_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.fixed_version_id",
			mcp.Description("The expression of issue fixed version id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.category_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.subject",
			mcp.Description("The expression of issue subject of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.user_id",
			mcp.Description("The expression of user id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.group",
			mcp.Description("The expression of user group of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.role",
			mcp.Description("The expression of user role of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.activity_id",
			mcp.Description("The expression of activity id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.comments",
			mcp.Description("The expression of comments of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.hours",
			mcp.Description("The expression of hours of the time entry. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogIndexProjectHandler)
}

func timelogIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseTimelogIndexProject(request)
	return toResult(c.TimelogIndexProject(ctx, project_id, &params, authorizationHeader))
}

func parseTimelogIndexProject(request mcp.CallToolRequest) client.TimelogIndexProjectParams {
	params := client.TimelogIndexProjectParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseTimelogIndexProjectQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func parseTimelogIndexProjectQuery(request *mcp.CallToolRequest) *client.TimelogIndexProjectParams_Query {
	params := client.TimelogIndexProjectParams_Query{}

	spent_on := request.GetString("query.spent_on", "")
	if spent_on != "" {

		params.SpentOn = &spent_on
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	issue_tracker_id := request.GetString("query.issue.tracker_id", "")
	if issue_tracker_id != "" {

		params.IssueTrackerId = &issue_tracker_id
	}

	issue_parent_id := request.GetString("query.issue.parent_id", "")
	if issue_parent_id != "" {

		params.IssueParentId = &issue_parent_id
	}

	issue_status_id := request.GetString("query.issue.status_id", "")
	if issue_status_id != "" {

		params.IssueStatusId = &issue_status_id
	}

	issue_fixed_version_id := request.GetString("query.issue.fixed_version_id", "")
	if issue_fixed_version_id != "" {

		params.IssueFixedVersionId = &issue_fixed_version_id
	}

	issue_category_id := request.GetString("query.issue.category_id", "")
	if issue_category_id != "" {

		params.IssueCategoryId = &issue_category_id
	}

	issue_subject := request.GetString("query.issue.subject", "")
	if issue_subject != "" {

		params.IssueSubject = &issue_subject
	}

	user_id := request.GetString("query.user_id", "")
	if user_id != "" {

		params.UserId = &user_id
	}

	user_group := request.GetString("query.user.group", "")
	if user_group != "" {

		params.UserGroup = &user_group
	}

	user_role := request.GetString("query.user.role", "")
	if user_role != "" {

		params.UserRole = &user_role
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	activity_id := request.GetString("query.activity_id", "")
	if activity_id != "" {

		params.ActivityId = &activity_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	issue_comments := request.GetString("query.issue.comments", "")
	if issue_comments != "" {

		params.IssueComments = &issue_comments
	}

	hours := request.GetString("query.hours", "")
	if hours != "" {

		params.Hours = &hours
	}

	return &params
}

func registerTimelogIndexCsv(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_index_csv",
		mcp.WithDescription("Returns a list of time entries."),
		mcp.WithString("query.spent_on",
			mcp.Description("The expression of spent_on of the time entry. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.tracker_id",
			mcp.Description("The expression of issue tracker id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.parent_id",
			mcp.Description("The expression of issue parent id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.status_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.fixed_version_id",
			mcp.Description("The expression of issue fixed version id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.category_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.subject",
			mcp.Description("The expression of issue subject of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.user_id",
			mcp.Description("The expression of user id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.group",
			mcp.Description("The expression of user group of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.role",
			mcp.Description("The expression of user role of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.activity_id",
			mcp.Description("The expression of activity id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.comments",
			mcp.Description("The expression of comments of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.hours",
			mcp.Description("The expression of hours of the time entry. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogIndexCsvHandler)
}

func timelogIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseTimelogIndexCsv(request)
	return toResult(c.TimelogIndexCsv(ctx, &params, authorizationHeader))
}

func parseTimelogIndexCsv(request mcp.CallToolRequest) client.TimelogIndexCsvParams {
	params := client.TimelogIndexCsvParams{}

	params.Query = parseTimelogIndexCsvQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func parseTimelogIndexCsvQuery(request *mcp.CallToolRequest) *client.TimelogIndexCsvParams_Query {
	params := client.TimelogIndexCsvParams_Query{}

	spent_on := request.GetString("query.spent_on", "")
	if spent_on != "" {

		params.SpentOn = &spent_on
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	issue_tracker_id := request.GetString("query.issue.tracker_id", "")
	if issue_tracker_id != "" {

		params.IssueTrackerId = &issue_tracker_id
	}

	issue_parent_id := request.GetString("query.issue.parent_id", "")
	if issue_parent_id != "" {

		params.IssueParentId = &issue_parent_id
	}

	issue_status_id := request.GetString("query.issue.status_id", "")
	if issue_status_id != "" {

		params.IssueStatusId = &issue_status_id
	}

	issue_fixed_version_id := request.GetString("query.issue.fixed_version_id", "")
	if issue_fixed_version_id != "" {

		params.IssueFixedVersionId = &issue_fixed_version_id
	}

	issue_category_id := request.GetString("query.issue.category_id", "")
	if issue_category_id != "" {

		params.IssueCategoryId = &issue_category_id
	}

	issue_subject := request.GetString("query.issue.subject", "")
	if issue_subject != "" {

		params.IssueSubject = &issue_subject
	}

	user_id := request.GetString("query.user_id", "")
	if user_id != "" {

		params.UserId = &user_id
	}

	user_group := request.GetString("query.user.group", "")
	if user_group != "" {

		params.UserGroup = &user_group
	}

	user_role := request.GetString("query.user.role", "")
	if user_role != "" {

		params.UserRole = &user_role
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	activity_id := request.GetString("query.activity_id", "")
	if activity_id != "" {

		params.ActivityId = &activity_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	issue_comments := request.GetString("query.issue.comments", "")
	if issue_comments != "" {

		params.IssueComments = &issue_comments
	}

	hours := request.GetString("query.hours", "")
	if hours != "" {

		params.Hours = &hours
	}

	return &params
}

func registerTimelogIndex(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_index",
		mcp.WithDescription("Returns a list of time entries."),
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
		mcp.WithString("query.spent_on",
			mcp.Description("The expression of spent_on of the time entry. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.tracker_id",
			mcp.Description("The expression of issue tracker id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.parent_id",
			mcp.Description("The expression of issue parent id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.status_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.fixed_version_id",
			mcp.Description("The expression of issue fixed version id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.category_id",
			mcp.Description("The expression of issue status id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.subject",
			mcp.Description("The expression of issue subject of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.user_id",
			mcp.Description("The expression of user id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.group",
			mcp.Description("The expression of user group of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.user.role",
			mcp.Description("The expression of user role of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.activity_id",
			mcp.Description("The expression of activity id of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the time entry. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.issue.comments",
			mcp.Description("The expression of comments of the time entry. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.hours",
			mcp.Description("The expression of hours of the time entry. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogIndexHandler)
}

func timelogIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseTimelogIndex(request)
	return toResult(c.TimelogIndex(ctx, &params, authorizationHeader))
}

func parseTimelogIndex(request mcp.CallToolRequest) client.TimelogIndexParams {
	params := client.TimelogIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseTimelogIndexQuery(&request)

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func parseTimelogIndexQuery(request *mcp.CallToolRequest) *client.TimelogIndexParams_Query {
	params := client.TimelogIndexParams_Query{}

	spent_on := request.GetString("query.spent_on", "")
	if spent_on != "" {

		params.SpentOn = &spent_on
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	issue_tracker_id := request.GetString("query.issue.tracker_id", "")
	if issue_tracker_id != "" {

		params.IssueTrackerId = &issue_tracker_id
	}

	issue_parent_id := request.GetString("query.issue.parent_id", "")
	if issue_parent_id != "" {

		params.IssueParentId = &issue_parent_id
	}

	issue_status_id := request.GetString("query.issue.status_id", "")
	if issue_status_id != "" {

		params.IssueStatusId = &issue_status_id
	}

	issue_fixed_version_id := request.GetString("query.issue.fixed_version_id", "")
	if issue_fixed_version_id != "" {

		params.IssueFixedVersionId = &issue_fixed_version_id
	}

	issue_category_id := request.GetString("query.issue.category_id", "")
	if issue_category_id != "" {

		params.IssueCategoryId = &issue_category_id
	}

	issue_subject := request.GetString("query.issue.subject", "")
	if issue_subject != "" {

		params.IssueSubject = &issue_subject
	}

	user_id := request.GetString("query.user_id", "")
	if user_id != "" {

		params.UserId = &user_id
	}

	user_group := request.GetString("query.user.group", "")
	if user_group != "" {

		params.UserGroup = &user_group
	}

	user_role := request.GetString("query.user.role", "")
	if user_role != "" {

		params.UserRole = &user_role
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	activity_id := request.GetString("query.activity_id", "")
	if activity_id != "" {

		params.ActivityId = &activity_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	issue_comments := request.GetString("query.issue.comments", "")
	if issue_comments != "" {

		params.IssueComments = &issue_comments
	}

	hours := request.GetString("query.hours", "")
	if hours != "" {

		params.Hours = &hours
	}

	return &params
}

func registerTimelogDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_destroy",
		mcp.WithDescription("Deletes the time entry with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the time entry."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogDestroyHandler)
}

func timelogDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseTimelogDestroy(request)
	return toResult(c.TimelogDestroy(ctx, id, &params, authorizationHeader))
}

func parseTimelogDestroy(request mcp.CallToolRequest) client.TimelogDestroyParams {
	params := client.TimelogDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerTimelogShow(s *server.MCPServer) {
	tool := mcp.NewTool("timelog_show",
		mcp.WithDescription("Returns the time entry with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the time entry."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, timelogShowHandler)
}

func timelogShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseTimelogShow(request)
	return toResult(c.TimelogShow(ctx, id, &params, authorizationHeader))
}

func parseTimelogShow(request mcp.CallToolRequest) client.TimelogShowParams {
	params := client.TimelogShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
