package redmine

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerCustomFieldsIndex(s *server.MCPServer) {
	tool := mcp.NewTool("custom_fields_index",
		mcp.WithDescription("Returns all custom field definitions."),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, customFieldsIndexHandler)
}

func customFieldsIndexHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	params := parseCustomFieldsIndex(request)
	return toResult(c.CustomFieldsIndex(ctx, &params, authorizationHeader))
}

func parseCustomFieldsIndex(request mcp.CallToolRequest) client.CustomFieldsIndexParams {
	params := client.CustomFieldsIndexParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
