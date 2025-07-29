package redmine

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerEnumerationsIndexDocumentCategory(s *server.MCPServer) {
	tool := mcp.NewTool("enumerations_index_document_category",
		mcp.WithDescription("Returns a list of document categories."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, enumerationsIndexDocumentCategoryHandler)
}

func enumerationsIndexDocumentCategoryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseEnumerationsIndexDocumentCategory(request)
	return toResult(c.EnumerationsIndexDocumentCategory(ctx, &params, authorizationHeader))
}

func parseEnumerationsIndexDocumentCategory(request mcp.CallToolRequest) client.EnumerationsIndexDocumentCategoryParams {
	params := client.EnumerationsIndexDocumentCategoryParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerEnumerationsIndexIssuePriority(s *server.MCPServer) {
	tool := mcp.NewTool("enumerations_index_issue_priority",
		mcp.WithDescription("Returns a list of issue priorities."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, enumerationsIndexIssuePriorityHandler)
}

func enumerationsIndexIssuePriorityHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseEnumerationsIndexIssuePriority(request)
	return toResult(c.EnumerationsIndexIssuePriority(ctx, &params, authorizationHeader))
}

func parseEnumerationsIndexIssuePriority(request mcp.CallToolRequest) client.EnumerationsIndexIssuePriorityParams {
	params := client.EnumerationsIndexIssuePriorityParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerEnumerationsIndexTimeEntryActivity(s *server.MCPServer) {
	tool := mcp.NewTool("enumerations_index_time_entry_activity",
		mcp.WithDescription("Returns a list of time entry activities."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, enumerationsIndexTimeEntryActivityHandler)
}

func enumerationsIndexTimeEntryActivityHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseEnumerationsIndexTimeEntryActivity(request)
	return toResult(c.EnumerationsIndexTimeEntryActivity(ctx, &params, authorizationHeader))
}

func parseEnumerationsIndexTimeEntryActivity(request mcp.CallToolRequest) client.EnumerationsIndexTimeEntryActivityParams {
	params := client.EnumerationsIndexTimeEntryActivityParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
