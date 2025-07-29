package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerIssuesIndexCsv(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index_csv",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexCsvHandler)
}

func issuesIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesIndexCsv(request)
	return toResult(c.IssuesIndexCsv(ctx, &params, authorizationHeader))
}

func parseIssuesIndexCsv(request mcp.CallToolRequest) client.IssuesIndexCsvParams {
	params := client.IssuesIndexCsvParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexCsvQuery(&request)

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

func parseIssuesIndexCsvQuery(request *mcp.CallToolRequest) *client.IssuesIndexCsvParams_Query {
	params := client.IssuesIndexCsvParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}

func registerIssuesIndex(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexHandler)
}

func issuesIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesIndex(request)
	return toResult(c.IssuesIndex(ctx, &params, authorizationHeader))
}

func parseIssuesIndex(request mcp.CallToolRequest) client.IssuesIndexParams {
	params := client.IssuesIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexQuery(&request)

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

func parseIssuesIndexQuery(request *mcp.CallToolRequest) *client.IssuesIndexParams_Query {
	params := client.IssuesIndexParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}

func registerIssuesIndexPdf(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index_pdf",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexPdfHandler)
}

func issuesIndexPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseIssuesIndexPdf(request)
	return toResult(c.IssuesIndexPdf(ctx, &params, authorizationHeader))
}

func parseIssuesIndexPdf(request mcp.CallToolRequest) client.IssuesIndexPdfParams {
	params := client.IssuesIndexPdfParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexPdfQuery(&request)

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

func parseIssuesIndexPdfQuery(request *mcp.CallToolRequest) *client.IssuesIndexPdfParams_Query {
	params := client.IssuesIndexPdfParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}

func registerIssuesDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("issues_destroy",
		mcp.WithDescription("Deletes the issue with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesDestroyHandler)
}

func issuesDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssuesDestroy(request)
	return toResult(c.IssuesDestroy(ctx, id, &params, authorizationHeader))
}

func parseIssuesDestroy(request mcp.CallToolRequest) client.IssuesDestroyParams {
	params := client.IssuesDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerIssuesShow(s *server.MCPServer) {
	tool := mcp.NewTool("issues_show",
		mcp.WithDescription("Returns the issue with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesShowHandler)
}

func issuesShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssuesShow(request)
	return toResult(c.IssuesShow(ctx, id, &params, authorizationHeader))
}

func parseIssuesShow(request mcp.CallToolRequest) client.IssuesShowParams {
	params := client.IssuesShowParams{}

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

func registerIssuesShowPdf(s *server.MCPServer) {
	tool := mcp.NewTool("issues_show_pdf",
		mcp.WithDescription("Returns the issue with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the issue."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesShowPdfHandler)
}

func issuesShowPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseIssuesShowPdf(request)
	return toResult(c.IssuesShowPdf(ctx, id, &params, authorizationHeader))
}

func parseIssuesShowPdf(request mcp.CallToolRequest) client.IssuesShowPdfParams {
	params := client.IssuesShowPdfParams{}

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

func registerIssuesIndexProjectCsv(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index_project_csv",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexProjectCsvHandler)
}

func issuesIndexProjectCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseIssuesIndexProjectCsv(request)
	return toResult(c.IssuesIndexProjectCsv(ctx, project_id, &params, authorizationHeader))
}

func parseIssuesIndexProjectCsv(request mcp.CallToolRequest) client.IssuesIndexProjectCsvParams {
	params := client.IssuesIndexProjectCsvParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexProjectCsvQuery(&request)

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

func parseIssuesIndexProjectCsvQuery(request *mcp.CallToolRequest) *client.IssuesIndexProjectCsvParams_Query {
	params := client.IssuesIndexProjectCsvParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}

func registerIssuesIndexProject(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index_project",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexProjectHandler)
}

func issuesIndexProjectHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseIssuesIndexProject(request)
	return toResult(c.IssuesIndexProject(ctx, project_id, &params, authorizationHeader))
}

func parseIssuesIndexProject(request mcp.CallToolRequest) client.IssuesIndexProjectParams {
	params := client.IssuesIndexProjectParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexProjectQuery(&request)

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

func parseIssuesIndexProjectQuery(request *mcp.CallToolRequest) *client.IssuesIndexProjectParams_Query {
	params := client.IssuesIndexProjectParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}

func registerIssuesIndexProjectPdf(s *server.MCPServer) {
	tool := mcp.NewTool("issues_index_project_pdf",
		mcp.WithDescription("Returns a paginated list of issues. By default, it returns open issues only."),
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
		mcp.WithString("query.status_id",
			mcp.Description("The expression of status id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project_id",
			mcp.Description("The expression of project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.tracker_id",
			mcp.Description("The expression of tracker id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.priority_id",
			mcp.Description("The expression of priority id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author_id",
			mcp.Description("The expression of author id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.group",
			mcp.Description("The expression of author group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.author.role",
			mcp.Description("The expression of author role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_id",
			mcp.Description("The expression of assigned to id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.member_of_group",
			mcp.Description("The expression of member of group of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.assigned_to_role",
			mcp.Description("The expression of assigned to role of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version_id",
			mcp.Description("The expression of fixed version id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.fixed_version.due_date",
			mcp.Description("The expression of fixed version due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.fixed_version.status",
			mcp.Description("The expression of fixed version status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.category_id",
			mcp.Description("The expression of category id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subject",
			mcp.Description("The expression of subject of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.description",
			mcp.Description("The expression of description of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.notes",
			mcp.Description("The expression of notes of the issue. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.updated_on",
			mcp.Description("The expression of updated_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.closed_on",
			mcp.Description("The expression of closed_on of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.start_date",
			mcp.Description("The expression of start date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.due_date",
			mcp.Description("The expression of due date of the issue. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.estimated_hours",
			mcp.Description("The expression of estimated hours of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.spent_time",
			mcp.Description("The expression of spent time of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.done_ratio",
			mcp.Description("The expression of done ratio of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.is_private",
			mcp.Description("The expression of private id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.attachment",
			mcp.Description("The expression of attachment of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.attachment_description",
			mcp.Description("The expression of attachment description of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.watcher_id",
			mcp.Description("The expression of watcher id of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.updated_by",
			mcp.Description("The expression of updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.last_updated_by",
			mcp.Description("The expression of last updated by of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.subproject_id",
			mcp.Description("The expression of sub project ID or identifier of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.project.status",
			mcp.Description("The expression of project status of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.relation_type",
			mcp.Description("The expression of relation type of the issue. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.parent_id",
			mcp.Description("The expression of parent id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.child_id",
			mcp.Description("The expression of child id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.issue_id",
			mcp.Description("The expression of issue id of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("query.any_searchable",
			mcp.Description("The expression of any searchable of the issue. The expression format is \"[operator]\\<value>\"."),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `allowed_statuses`, `attachments`, `changesets`, `children`, `journals`, `relations`, `watchers`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, issuesIndexProjectPdfHandler)
}

func issuesIndexProjectPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseIssuesIndexProjectPdf(request)
	return toResult(c.IssuesIndexProjectPdf(ctx, project_id, &params, authorizationHeader))
}

func parseIssuesIndexProjectPdf(request mcp.CallToolRequest) client.IssuesIndexProjectPdfParams {
	params := client.IssuesIndexProjectPdfParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseIssuesIndexProjectPdfQuery(&request)

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

func parseIssuesIndexProjectPdfQuery(request *mcp.CallToolRequest) *client.IssuesIndexProjectPdfParams_Query {
	params := client.IssuesIndexProjectPdfParams_Query{}

	status_id := request.GetString("query.status_id", "")
	if status_id != "" {

		params.StatusId = &status_id
	}

	project_id := request.GetString("query.project_id", "")
	if project_id != "" {

		params.ProjectId = &project_id
	}

	tracker_id := request.GetString("query.tracker_id", "")
	if tracker_id != "" {

		params.TrackerId = &tracker_id
	}

	priority_id := request.GetString("query.priority_id", "")
	if priority_id != "" {

		params.PriorityId = &priority_id
	}

	author_id := request.GetString("query.author_id", "")
	if author_id != "" {

		params.AuthorId = &author_id
	}

	author_group := request.GetString("query.author.group", "")
	if author_group != "" {

		params.AuthorGroup = &author_group
	}

	author_role := request.GetString("query.author.role", "")
	if author_role != "" {

		params.AuthorRole = &author_role
	}

	assigned_to_id := request.GetString("query.assigned_to_id", "")
	if assigned_to_id != "" {

		params.AssignedToId = &assigned_to_id
	}

	member_of_group := request.GetString("query.member_of_group", "")
	if member_of_group != "" {

		params.MemberOfGroup = &member_of_group
	}

	assigned_to_role := request.GetString("query.assigned_to_role", "")
	if assigned_to_role != "" {

		params.AssignedToRole = &assigned_to_role
	}

	fixed_version_id := request.GetString("query.fixed_version_id", "")
	if fixed_version_id != "" {

		params.FixedVersionId = &fixed_version_id
	}

	fixed_version_due_date := request.GetString("query.fixed_version.due_date", "")
	if fixed_version_due_date != "" {

		params.FixedVersionDueDate = &fixed_version_due_date
	}

	fixed_version_status := request.GetString("query.fixed_version.status", "")
	if fixed_version_status != "" {

		params.FixedVersionStatus = &fixed_version_status
	}

	category_id := request.GetString("query.category_id", "")
	if category_id != "" {

		params.CategoryId = &category_id
	}

	subject := request.GetString("query.subject", "")
	if subject != "" {

		params.Subject = &subject
	}

	description := request.GetString("query.description", "")
	if description != "" {

		params.Description = &description
	}

	notes := request.GetString("query.notes", "")
	if notes != "" {

		params.Notes = &notes
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	updated_on := request.GetString("query.updated_on", "")
	if updated_on != "" {

		params.UpdatedOn = &updated_on
	}

	closed_on := request.GetString("query.closed_on", "")
	if closed_on != "" {

		params.ClosedOn = &closed_on
	}

	start_date := request.GetString("query.start_date", "")
	if start_date != "" {

		params.StartDate = &start_date
	}

	due_date := request.GetString("query.due_date", "")
	if due_date != "" {

		params.DueDate = &due_date
	}

	estimated_hours := request.GetString("query.estimated_hours", "")
	if estimated_hours != "" {

		params.EstimatedHours = &estimated_hours
	}

	spent_time := request.GetString("query.spent_time", "")
	if spent_time != "" {

		params.SpentTime = &spent_time
	}

	done_ratio := request.GetString("query.done_ratio", "")
	if done_ratio != "" {

		params.DoneRatio = &done_ratio
	}

	is_private := request.GetString("query.is_private", "")
	if is_private != "" {

		params.IsPrivate = &is_private
	}

	attachment := request.GetString("query.attachment", "")
	if attachment != "" {

		params.Attachment = &attachment
	}

	attachment_description := request.GetString("query.attachment_description", "")
	if attachment_description != "" {

		params.AttachmentDescription = &attachment_description
	}

	watcher_id := request.GetString("query.watcher_id", "")
	if watcher_id != "" {

		params.WatcherId = &watcher_id
	}

	updated_by := request.GetString("query.updated_by", "")
	if updated_by != "" {

		params.UpdatedBy = &updated_by
	}

	last_updated_by := request.GetString("query.last_updated_by", "")
	if last_updated_by != "" {

		params.LastUpdatedBy = &last_updated_by
	}

	subproject_id := request.GetString("query.subproject_id", "")
	if subproject_id != "" {

		params.SubprojectId = &subproject_id
	}

	project_status := request.GetString("query.project.status", "")
	if project_status != "" {

		params.ProjectStatus = &project_status
	}

	relation_type := request.GetString("query.relation_type", "")
	if relation_type != "" {

		params.RelationType = &relation_type
	}

	parent_id := request.GetString("query.parent_id", "")
	if parent_id != "" {

		params.ParentId = &parent_id
	}

	child_id := request.GetString("query.child_id", "")
	if child_id != "" {

		params.ChildId = &child_id
	}

	issue_id := request.GetString("query.issue_id", "")
	if issue_id != "" {

		params.IssueId = &issue_id
	}

	any_searchable := request.GetString("query.any_searchable", "")
	if any_searchable != "" {

		params.AnySearchable = &any_searchable
	}

	return &params
}
