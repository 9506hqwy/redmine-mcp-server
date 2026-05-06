package redmine

import (
	"context"
	"encoding/json"

	"github.com/invopop/jsonschema"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

type AttachmentsDownloadRequest struct {
	Id     int                               `json:"id" jsonschema:"description=The ID of the attachment."`
	Params *client.AttachmentsDownloadParams `json:"params,omitempty"`
}

func registerAttachmentsDownload(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsDownloadRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_download",
		mcp.WithDescription("Download the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsDownloadHandler))
}

func attachmentsDownloadHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsDownloadRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsDownload(ctx, req.Id, req.Params, authorizationHeader))
}

type AttachmentsThumbnailRequest struct {
	Id     int                                `json:"id" jsonschema:"description=The ID of the attachment."`
	Params *client.AttachmentsThumbnailParams `json:"params,omitempty"`
}

func registerAttachmentsThumbnail(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsThumbnailRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_thumbnail",
		mcp.WithDescription("Download the thumbnail of the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsThumbnailHandler))
}

func attachmentsThumbnailHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsThumbnailRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsThumbnail(ctx, req.Id, req.Params, authorizationHeader))
}

type AttachmentsThumbnailSizeRequest struct {
	Id     int                                    `json:"id" jsonschema:"description=The ID of the attachment."`
	Size   int                                    `json:"size" jsonschema:"description=The size of the attachment."`
	Params *client.AttachmentsThumbnailSizeParams `json:"params,omitempty"`
}

func registerAttachmentsThumbnailSize(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsThumbnailSizeRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_thumbnail_size",
		mcp.WithDescription("Downloads the attachment thumbnail with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsThumbnailSizeHandler))
}

func attachmentsThumbnailSizeHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsThumbnailSizeRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsThumbnailSize(ctx, req.Id, req.Size, req.Params, authorizationHeader))
}

type AttachmentsDestroyRequest struct {
	Id     int                              `json:"id" jsonschema:"description=The ID of the attachment."`
	Params *client.AttachmentsDestroyParams `json:"params,omitempty"`
}

func registerAttachmentsDestroy(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsDestroyRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_destroy",
		mcp.WithDescription("Deletes the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsDestroyHandler))
}

func attachmentsDestroyHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsDestroyRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsDestroy(ctx, req.Id, req.Params, authorizationHeader))
}

type AttachmentsUpdatePatchRequest struct {
	Id     int                                          `json:"id" jsonschema:"description=The ID of the attachment."`
	Params *client.AttachmentsUpdatePatchParams         `json:"params,omitempty"`
	Body   client.AttachmentsUpdatePatchJSONRequestBody `json:"body,omitempty"`
}

func registerAttachmentsUpdatePatch(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsUpdatePatchRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_update_patch",
		mcp.WithDescription("Updates the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsUpdatePatchHandler))
}

func attachmentsUpdatePatchHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsUpdatePatchRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsUpdatePatch(ctx, req.Id, req.Params, req.Body, authorizationHeader))
}

type AttachmentsShowRequest struct {
	Id     int                           `json:"id" jsonschema:"description=The ID of the attachment."`
	Params *client.AttachmentsShowParams `json:"params,omitempty"`
}

func registerAttachmentsShow(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsShowRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_show",
		mcp.WithDescription("Returns the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsShowHandler))
}

func attachmentsShowHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsShowRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsShow(ctx, req.Id, req.Params, authorizationHeader))
}

type AttachmentsDownloadAllRequest struct {
	ObjectType string                               `json:"object_type" jsonschema:"description=The object type of the attachment."`
	ObjectId   int                                  `json:"object_id" jsonschema:"description=The object ID of the attachment."`
	Params     *client.AttachmentsDownloadAllParams `json:"params,omitempty"`
}

func registerAttachmentsDownloadAll(s *server.MCPServer) {
	r := &jsonschema.Reflector{}
	r.DoNotReference = true
	schemaObj := r.Reflect(&AttachmentsDownloadAllRequest{})
	mcpSchema, err := json.Marshal(schemaObj)
	if err != nil {
		return
	}

	rawSchema := json.RawMessage(mcpSchema)

	tool := mcp.NewTool("attachments_download_all",
		mcp.WithDescription("Downloads the attachment with the specified ID."),
		mcp.WithRawInputSchema(rawSchema),
		func(tool *mcp.Tool) {
			tool.InputSchema.Type = ""
		},
	)

	s.AddTool(tool, mcp.NewTypedToolHandler(attachmentsDownloadAllHandler))
}

func attachmentsDownloadAllHandler(ctx context.Context, request mcp.CallToolRequest, req AttachmentsDownloadAllRequest) (*mcp.CallToolResult, error) {
	c, err := newClient(ctx)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return toResult(c.AttachmentsDownloadAll(ctx, req.ObjectType, req.ObjectId, req.Params, authorizationHeader))
}
