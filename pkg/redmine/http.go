package redmine

import (
	"context"
	"fmt"
	"io"
	"net/http"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
	"github.com/mark3labs/mcp-go/mcp"
)

type ApiKeyKey struct{}
type UrlKey struct{}
type UserKey struct{}
type PasswordKey struct{}

func authorizationHeader(ctx context.Context, req *http.Request) error {
	if apiKeyAuth(ctx, req) == nil {
		return nil
	}

	return basicAuth(ctx, req)
}

func basicAuth(ctx context.Context, req *http.Request) error {
	user, ok := ctx.Value(UserKey{}).(string)
	if !ok || user == "" {
		return fmt.Errorf("missing user")
	}

	password, ok := ctx.Value(PasswordKey{}).(string)
	if !ok {
		return fmt.Errorf("missing password")
	}

	req.SetBasicAuth(user, password)
	return nil
}

func apiKeyAuth(ctx context.Context, req *http.Request) error {
	apiKey, ok := ctx.Value(ApiKeyKey{}).(string)
	if !ok || apiKey == "" {
		return fmt.Errorf("missing API key")
	}

	req.Header.Set("X-Redmine-API-Key", apiKey)
	return nil
}

func newClient(ctx context.Context) (*client.ClientWithResponses, error) {
	hc := http.Client{}

	url, ok := ctx.Value(UrlKey{}).(string)
	if !ok || url == "" {
		return nil, fmt.Errorf("missing url")
	}

	return client.NewClientWithResponses(url, client.WithHTTPClient(&hc))
}

func toResult(response *http.Response, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if response.StatusCode < http.StatusOK || http.StatusMultipleChoices <= response.StatusCode {
		return mcp.NewToolResultError(fmt.Sprintf("%s: %s", response.Status, string(body))), nil
	}

	return mcp.NewToolResultText(string(body)), nil
}
