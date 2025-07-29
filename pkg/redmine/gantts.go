package redmine

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerGanttsShowPdf(s *server.MCPServer) {
	tool := mcp.NewTool("gantts_show_pdf",
		mcp.WithDescription("Download the Gantt chart."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, ganttsShowPdfHandler)
}

func ganttsShowPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGanttsShowPdf(request)
	return toResult(c.GanttsShowPdf(ctx, &params, authorizationHeader))
}

func parseGanttsShowPdf(request mcp.CallToolRequest) client.GanttsShowPdfParams {
	params := client.GanttsShowPdfParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGanttsShowPng(s *server.MCPServer) {
	tool := mcp.NewTool("gantts_show_png",
		mcp.WithDescription("Download the Gantt chart."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, ganttsShowPngHandler)
}

func ganttsShowPngHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseGanttsShowPng(request)
	return toResult(c.GanttsShowPng(ctx, &params, authorizationHeader))
}

func parseGanttsShowPng(request mcp.CallToolRequest) client.GanttsShowPngParams {
	params := client.GanttsShowPngParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGanttsShowProjectPdf(s *server.MCPServer) {
	tool := mcp.NewTool("gantts_show_project_pdf",
		mcp.WithDescription("Download the Gantt chart for the specified project."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, ganttsShowProjectPdfHandler)
}

func ganttsShowProjectPdfHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseGanttsShowProjectPdf(request)
	return toResult(c.GanttsShowProjectPdf(ctx, project_id, &params, authorizationHeader))
}

func parseGanttsShowProjectPdf(request mcp.CallToolRequest) client.GanttsShowProjectPdfParams {
	params := client.GanttsShowProjectPdfParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerGanttsShowProjectPng(s *server.MCPServer) {
	tool := mcp.NewTool("gantts_show_project_png",
		mcp.WithDescription("Download the Gantt chart for the specified project."),
		mcp.WithString("project_id",
			mcp.Description("The ID or identifier of the project."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, ganttsShowProjectPngHandler)
}

func ganttsShowProjectPngHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	project_id := request.GetString("project_id", "")
	params := parseGanttsShowProjectPng(request)
	return toResult(c.GanttsShowProjectPng(ctx, project_id, &params, authorizationHeader))
}

func parseGanttsShowProjectPng(request mcp.CallToolRequest) client.GanttsShowProjectPngParams {
	params := client.GanttsShowProjectPngParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
