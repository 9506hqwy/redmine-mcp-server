package redmine

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
	"github.com/mark3labs/mcp-go/mcp"
)

type APIKeyKey struct{}
type URKKey struct{}
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
		return errors.New("missing user")
	}

	password, ok := ctx.Value(PasswordKey{}).(string)
	if !ok {
		return errors.New("missing password")
	}

	req.SetBasicAuth(user, password)
	return nil
}

func apiKeyAuth(ctx context.Context, req *http.Request) error {
	apiKey, ok := ctx.Value(APIKeyKey{}).(string)
	if !ok || apiKey == "" {
		return errors.New("missing API key")
	}

	req.Header.Set("X-Redmine-API-Key", apiKey)
	return nil
}

func newClient(ctx context.Context) (*client.ClientWithResponses, error) {
	hc := http.Client{}

	url, ok := ctx.Value(URKKey{}).(string)
	if !ok || url == "" {
		return nil, errors.New("missing url")
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
