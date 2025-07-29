package redmine

import (
	"context"
	"math"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerWikiShowRoot(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_root",
		mcp.WithDescription("Returns the details of the root wiki page."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowRootHandler)
}

func wikiShowRootHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseWikiShowRoot(request)
	return toResult(c.WikiShowRoot(ctx, project_id, &params, authorizationHeader))
}

func parseWikiShowRoot(request mcp.CallToolRequest) client.WikiShowRootParams {
	params := client.WikiShowRootParams{}

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

func registerWikiIndex(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_index",
		mcp.WithDescription("Returns a list of all pages in the project wiki."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiIndexHandler)
}

func wikiIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseWikiIndex(request)
	return toResult(c.WikiIndex(ctx, project_id, &params, authorizationHeader))
}

func parseWikiIndex(request mcp.CallToolRequest) client.WikiIndexParams {
	params := client.WikiIndexParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerWikiDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_destroy",
		mcp.WithDescription("Deletes a wiki page, its attachments and its history with the specified ID. If the deleted page is a parent page, its child pages are not deleted but changed as root pages."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiDestroyHandler)
}

func wikiDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	params := parseWikiDestroy(request)
	return toResult(c.WikiDestroy(ctx, project_id, id, &params, authorizationHeader))
}

func parseWikiDestroy(request mcp.CallToolRequest) client.WikiDestroyParams {
	params := client.WikiDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerWikiShow(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowHandler)
}

func wikiShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	params := parseWikiShow(request)
	return toResult(c.WikiShow(ctx, project_id, id, &params, authorizationHeader))
}

func parseWikiShow(request mcp.CallToolRequest) client.WikiShowParams {
	params := client.WikiShowParams{}

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

func registerWikiShowPdf(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_pdf",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowPdfHandler)
}

func wikiShowPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	params := parseWikiShowPdf(request)
	return toResult(c.WikiShowPdf(ctx, project_id, id, &params, authorizationHeader))
}

func parseWikiShowPdf(request mcp.CallToolRequest) client.WikiShowPdfParams {
	params := client.WikiShowPdfParams{}

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

func registerWikiShowTxt(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_txt",
		mcp.WithDescription("Returns the details of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowTxtHandler)
}

func wikiShowTxtHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	params := parseWikiShowTxt(request)
	return toResult(c.WikiShowTxt(ctx, project_id, id, &params, authorizationHeader))
}

func parseWikiShowTxt(request mcp.CallToolRequest) client.WikiShowTxtParams {
	params := client.WikiShowTxtParams{}

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

func registerWikiShowVersion(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_version",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithNumber("version",
			mcp.Description("The version of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowVersionHandler)
}

func wikiShowVersionHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	version := request.GetInt("version", math.MinInt)
	params := parseWikiShowVersion(request)
	return toResult(c.WikiShowVersion(ctx, project_id, id, version, &params, authorizationHeader))
}

func parseWikiShowVersion(request mcp.CallToolRequest) client.WikiShowVersionParams {
	params := client.WikiShowVersionParams{}

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

func registerWikiShowVersionPdf(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_version_pdf",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithNumber("version",
			mcp.Description("The version of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowVersionPdfHandler)
}

func wikiShowVersionPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	version := request.GetInt("version", math.MinInt)
	params := parseWikiShowVersionPdf(request)
	return toResult(c.WikiShowVersionPdf(ctx, project_id, id, version, &params, authorizationHeader))
}

func parseWikiShowVersionPdf(request mcp.CallToolRequest) client.WikiShowVersionPdfParams {
	params := client.WikiShowVersionPdfParams{}

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

func registerWikiShowVersionTxt(s *server.MCPServer) {
	tool := mcp.NewTool("wiki_show_version_txt",
		mcp.WithDescription("Returns the details of an old version of a wiki page with the specified ID."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("id",
			mcp.Description("The title of the wiki."),
			mcp.Required(),
		),
		mcp.WithNumber("version",
			mcp.Description("The version of the wiki."),
			mcp.Required(),
		),
		mcp.WithString("include",
			mcp.Description("fetch associated data (optional). Possible values: `attachments`."),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, wikiShowVersionTxtHandler)
}

func wikiShowVersionTxtHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	id := request.GetString("id", "")
	version := request.GetInt("version", math.MinInt)
	params := parseWikiShowVersionTxt(request)
	return toResult(c.WikiShowVersionTxt(ctx, project_id, id, version, &params, authorizationHeader))
}

func parseWikiShowVersionTxt(request mcp.CallToolRequest) client.WikiShowVersionTxtParams {
	params := client.WikiShowVersionTxtParams{}

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
