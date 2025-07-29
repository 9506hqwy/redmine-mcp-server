package redmine

import (
	"context"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func registerAttachmentsDownload(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_download",
		mcp.WithDescription("Download the attachment with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsDownloadHandler)
}

func attachmentsDownloadHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseAttachmentsDownload(request)
	return toResult(c.AttachmentsDownload(ctx, id, &params, authorizationHeader))
}

func parseAttachmentsDownload(request mcp.CallToolRequest) client.AttachmentsDownloadParams {
	params := client.AttachmentsDownloadParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerAttachmentsThumbnail(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_thumbnail",
		mcp.WithDescription("Download the thumbnail of the attachment with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsThumbnailHandler)
}

func attachmentsThumbnailHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseAttachmentsThumbnail(request)
	return toResult(c.AttachmentsThumbnail(ctx, id, &params, authorizationHeader))
}

func parseAttachmentsThumbnail(request mcp.CallToolRequest) client.AttachmentsThumbnailParams {
	params := client.AttachmentsThumbnailParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerAttachmentsThumbnailSize(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_thumbnail_size",
		mcp.WithDescription("Downloads the attachment thumbnail with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithNumber("size",
			mcp.Description("The size of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsThumbnailSizeHandler)
}

func attachmentsThumbnailSizeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	size := request.GetInt("size", math.MinInt)
	params := parseAttachmentsThumbnailSize(request)
	return toResult(c.AttachmentsThumbnailSize(ctx, id, size, &params, authorizationHeader))
}

func parseAttachmentsThumbnailSize(request mcp.CallToolRequest) client.AttachmentsThumbnailSizeParams {
	params := client.AttachmentsThumbnailSizeParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerAttachmentsDestroy(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_destroy",
		mcp.WithDescription("Deletes the attachment with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsDestroyHandler)
}

func attachmentsDestroyHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseAttachmentsDestroy(request)
	return toResult(c.AttachmentsDestroy(ctx, id, &params, authorizationHeader))
}

func parseAttachmentsDestroy(request mcp.CallToolRequest) client.AttachmentsDestroyParams {
	params := client.AttachmentsDestroyParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerAttachmentsShow(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_show",
		mcp.WithDescription("Returns the attachment with the specified ID."),
		mcp.WithNumber("id",
			mcp.Description("The ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsShowHandler)
}

func attachmentsShowHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	id := request.GetInt("id", math.MinInt)
	params := parseAttachmentsShow(request)
	return toResult(c.AttachmentsShow(ctx, id, &params, authorizationHeader))
}

func parseAttachmentsShow(request mcp.CallToolRequest) client.AttachmentsShowParams {
	params := client.AttachmentsShowParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}

func registerAttachmentsDownloadAll(s *server.MCPServer) {
	tool := mcp.NewTool("attachments_download_all",
		mcp.WithDescription("Downloads the attachment with the specified ID."),
		mcp.WithString("object_type",
			mcp.Description("The object type of the attachment."),
			mcp.Required(),
		),
		mcp.WithNumber("object_id",
			mcp.Description("The object ID of the attachment."),
			mcp.Required(),
		),
		mcp.WithString("X-Redmine-Switch-User",
			mcp.Description("This only works when using the API with an administrator account, this header will be ignored when using the API with a regular user account."),
		),
	)

	s.AddTool(tool, attachmentsDownloadAllHandler)
}

func attachmentsDownloadAllHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	object_type := request.GetString("object_type", "")
	object_id := request.GetInt("object_id", math.MinInt)
	params := parseAttachmentsDownloadAll(request)
	return toResult(c.AttachmentsDownloadAll(ctx, object_type, object_id, &params, authorizationHeader))
}

func parseAttachmentsDownloadAll(request mcp.CallToolRequest) client.AttachmentsDownloadAllParams {
	params := client.AttachmentsDownloadAllParams{}

	X_Redmine_Switch_User := request.GetString("X-Redmine-Switch-User", "")
	if X_Redmine_Switch_User != "" {

		params.XRedmineSwitchUser = &X_Redmine_Switch_User
	}

	return params
}
