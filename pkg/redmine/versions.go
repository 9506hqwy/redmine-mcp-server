package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerVersionsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("versions_index",
		mcp.WithDescription("Returns the versions available for the project with the specified ID or identifier (:project_id). The response may include shared versions from other projects."),
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

	s.AddTool(tool, versionsIndexHandler)
}

func versionsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseVersionsIndex(request)
	return toResult(c.VersionsIndex(ctx, project_id, &params, authorizationHeader))
}

func parseVersionsIndex(request mcp.CallToolRequest) client.VersionsIndexParams {
	params := client.VersionsIndexParams{}

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

func registerVersionsShowTxt(s *server.MCPServer) {
	tool := mcp.NewTool("versions_show_txt",
		mcp.WithDescription("Returns the version with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the version."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, versionsShowTxtHandler)
}

func versionsShowTxtHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseVersionsShowTxt(request)
	return toResult(c.VersionsShowTxt(ctx, id, &params, authorizationHeader))
}

func parseVersionsShowTxt(request mcp.CallToolRequest) client.VersionsShowTxtParams {
	params := client.VersionsShowTxtParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerVersionsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("versions_destroy",
		mcp.WithDescription("Deletes the version with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the version."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, versionsDestroyHandler)
}

func versionsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseVersionsDestroy(request)
	return toResult(c.VersionsDestroy(ctx, id, &params, authorizationHeader))
}

func parseVersionsDestroy(request mcp.CallToolRequest) client.VersionsDestroyParams {
	params := client.VersionsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerVersionsShow(s *server.MCPServer) {
	tool := mcp.NewTool("versions_show",
		mcp.WithDescription("Returns the version with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the version."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, versionsShowHandler)
}

func versionsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseVersionsShow(request)
	return toResult(c.VersionsShow(ctx, id, &params, authorizationHeader))
}

func parseVersionsShow(request mcp.CallToolRequest) client.VersionsShowParams {
	params := client.VersionsShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
