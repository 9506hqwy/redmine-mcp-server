package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerUsersIndexCsv(s *server.MCPServer) {
	tool := mcp.NewTool("users_index_csv",
		mcp.WithDescription("Returns a list of all users in CSV format."),
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
			mcp.Description("The expression of status of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `1`: active - `2`: registered - `3`: locked examples: - active or locked: `1|3` - not locked: `!3`"),
		),
		mcp.WithString("query.auth_source_id",
			mcp.Description("The expression of auth source id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.is_member_of_group",
			mcp.Description("The expression of group id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.twofa_scheme",
			mcp.Description("The expression of 2FA scheme id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.name",
			mcp.Description("The expression of name of the user. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.login",
			mcp.Description("The expression of login of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.firstname",
			mcp.Description("The expression of firstname of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.lastname",
			mcp.Description("The expression of lastname of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.mail",
			mcp.Description("The expression of mail of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the user. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.last_login_on",
			mcp.Description("The expression of last_login_on of the user. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.admin",
			mcp.Description("The expression of administrative of the user. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `0`: no - `1`: yes examples: - administrator: `1` - non administrator: `0`"),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `auth_source`, `memberships`, `groups`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, usersIndexCsvHandler)
}

func usersIndexCsvHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersIndexCsv(request)
	return toResult(c.UsersIndexCsv(ctx, &params, authorizationHeader))
}

func parseUsersIndexCsv(request mcp.CallToolRequest) client.UsersIndexCsvParams {
	params := client.UsersIndexCsvParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseUsersIndexCsvQuery(&request)

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

func parseUsersIndexCsvQuery(request *mcp.CallToolRequest) *client.UsersIndexCsvParams_Query {
	params := client.UsersIndexCsvParams_Query{}

	status := request.GetString("query.status", "")
	if status != "" {

		params.Status = &status
	}

	auth_source_id := request.GetString("query.auth_source_id", "")
	if auth_source_id != "" {

		params.AuthSourceId = &auth_source_id
	}

	is_member_of_group := request.GetString("query.is_member_of_group", "")
	if is_member_of_group != "" {

		params.IsMemberOfGroup = &is_member_of_group
	}

	twofa_scheme := request.GetString("query.twofa_scheme", "")
	if twofa_scheme != "" {

		params.TwofaScheme = &twofa_scheme
	}

	name := request.GetString("query.name", "")
	if name != "" {

		params.Name = &name
	}

	login := request.GetString("query.login", "")
	if login != "" {

		params.Login = &login
	}

	firstname := request.GetString("query.firstname", "")
	if firstname != "" {

		params.Firstname = &firstname
	}

	lastname := request.GetString("query.lastname", "")
	if lastname != "" {

		params.Lastname = &lastname
	}

	mail := request.GetString("query.mail", "")
	if mail != "" {

		params.Mail = &mail
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	last_login_on := request.GetString("query.last_login_on", "")
	if last_login_on != "" {

		params.LastLoginOn = &last_login_on
	}

	admin := request.GetString("query.admin", "")
	if admin != "" {

		params.Admin = &admin
	}

	return &params
}

func registerUsersIndex(s *server.MCPServer) {
	tool := mcp.NewTool("users_index",
		mcp.WithDescription("Returns a list of all users."),
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
			mcp.Description("The expression of status of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `1`: active - `2`: registered - `3`: locked examples: - active or locked: `1|3` - not locked: `!3`"),
		),
		mcp.WithString("query.auth_source_id",
			mcp.Description("The expression of auth source id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.is_member_of_group",
			mcp.Description("The expression of group id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.twofa_scheme",
			mcp.Description("The expression of 2FA scheme id of the user. The expression format is \"[operator]\\<values>\". Values should be separated by a pipe \"|\"."),
		),
		mcp.WithString("query.name",
			mcp.Description("The expression of name of the user. The expression format is \"\\<operator>\\<value>\". examples: - contains \"test\": `~test` - not contains \"test\": `!~test`"),
		),
		mcp.WithString("query.login",
			mcp.Description("The expression of login of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.firstname",
			mcp.Description("The expression of firstname of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.lastname",
			mcp.Description("The expression of lastname of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.mail",
			mcp.Description("The expression of mail of the user. The expression format is \"\\<operator>\\<value>\"."),
		),
		mcp.WithString("query.created_on",
			mcp.Description("The expression of created_on of the user. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.last_login_on",
			mcp.Description("The expression of last_login_on of the user. The expression format is \"[operator]\\<value>\". examples: - last week : `lw` - before \"2025/7/12\": `<=2025-07-12`"),
		),
		mcp.WithString("query.admin",
			mcp.Description("The expression of administrative of the user. The expression format is \"\\<operator>\\<values>\". Values should be separated by a pipe \"|\". Possible values are: - `0`: no - `1`: yes examples: - administrator: `1` - non administrator: `0`"),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `auth_source`, `memberships`, `groups`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, usersIndexHandler)
}

func usersIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseUsersIndex(request)
	return toResult(c.UsersIndex(ctx, &params, authorizationHeader))
}

func parseUsersIndex(request mcp.CallToolRequest) client.UsersIndexParams {
	params := client.UsersIndexParams{}

	X_Redmine_Nometa := request.GetInt("X-Redmine-Nometa", math.MinInt)
	if X_Redmine_Nometa != math.MinInt {

		params.XRedmineNometa = &X_Redmine_Nometa
	}

	params.Pagination = parsePagination(&request)

	params.Query = parseUsersIndexQuery(&request)

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

func parseUsersIndexQuery(request *mcp.CallToolRequest) *client.UsersIndexParams_Query {
	params := client.UsersIndexParams_Query{}

	status := request.GetString("query.status", "")
	if status != "" {

		params.Status = &status
	}

	auth_source_id := request.GetString("query.auth_source_id", "")
	if auth_source_id != "" {

		params.AuthSourceId = &auth_source_id
	}

	is_member_of_group := request.GetString("query.is_member_of_group", "")
	if is_member_of_group != "" {

		params.IsMemberOfGroup = &is_member_of_group
	}

	twofa_scheme := request.GetString("query.twofa_scheme", "")
	if twofa_scheme != "" {

		params.TwofaScheme = &twofa_scheme
	}

	name := request.GetString("query.name", "")
	if name != "" {

		params.Name = &name
	}

	login := request.GetString("query.login", "")
	if login != "" {

		params.Login = &login
	}

	firstname := request.GetString("query.firstname", "")
	if firstname != "" {

		params.Firstname = &firstname
	}

	lastname := request.GetString("query.lastname", "")
	if lastname != "" {

		params.Lastname = &lastname
	}

	mail := request.GetString("query.mail", "")
	if mail != "" {

		params.Mail = &mail
	}

	created_on := request.GetString("query.created_on", "")
	if created_on != "" {

		params.CreatedOn = &created_on
	}

	last_login_on := request.GetString("query.last_login_on", "")
	if last_login_on != "" {

		params.LastLoginOn = &last_login_on
	}

	admin := request.GetString("query.admin", "")
	if admin != "" {

		params.Admin = &admin
	}

	return &params
}

func registerUsersDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("users_destroy",
		mcp.WithDescription("Deletes the user with the specified ID."),
		mcp.WithString("id",
			mcp.Description("The ID or `current` of the user."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, usersDestroyHandler)
}

func usersDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseUsersDestroy(request)
	return toResult(c.UsersDestroy(ctx, id, &params, authorizationHeader))
}

func parseUsersDestroy(request mcp.CallToolRequest) client.UsersDestroyParams {
	params := client.UsersDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerUsersShow(s *server.MCPServer) {
	tool := mcp.NewTool("users_show",
		mcp.WithDescription("Returns the user with the specified ID. Use /users/current.json to retrieve the user whose credentials is used to access the API."),
		mcp.WithString("id",
			mcp.Description("The ID or `current` of the user."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `auth_source`, `memberships`, `groups`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, usersShowHandler)
}

func usersShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetString("id", "")
	params := parseUsersShow(request)
	return toResult(c.UsersShow(ctx, id, &params, authorizationHeader))
}

func parseUsersShow(request mcp.CallToolRequest) client.UsersShowParams {
	params := client.UsersShowParams{}

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
