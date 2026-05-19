package github

import (
	"context"

	"github.com/imroc/req/v3"
	"go.opentelemetry.io/otel/trace"
)

// Client is the go client for GitHub API.
type Client struct {
	*req.Client
}

// APIError represents the error message that GitHub API returns.
// GitHub API doc: https://docs.github.com/en/rest/overview/resources-in-the-rest-api#client-errors
type APIError struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url,omitempty"`
	Errors           []struct {
		Resource string `json:"resource"`
		Field    string `json:"field"`
		Code     string `json:"code"`
	} `json:"errors,omitempty"`
}

// Error convert APIError to a human readable error and return.
func (e *APIError) Error() string { _ = "STUB: not implemented"; return "" }

// NewClient create a GitHub client.
func NewClient() *Client {
	_ = "STUB: not implemented"

	// All GitHub API requests need this header.
	return nil
}

// All GitHub API requests use the same base URL.

// Enable dump at the request-level for each request, and only
// temporarily stores the dump content in memory, so we can call
// resp.Dump() to get the dump content when needed in response
// middleware.
// This is actually a syntax sugar, implemented internally using
// request middleware

// Unmarshal response body into an APIError struct when status >= 400.

// Handle common exceptions in response middleware.

// There is an underlying error, e.g. network error or unmarshal error (SetSuccessResult or SetErrorResult was invoked before).
// Append dump content to original underlying error to help troubleshoot.

// Skip the following logic if there is an underlying error.

// Server returns an error message.
// Convert it to human-readable go error which implements the error interface.

// Corner case: neither an error response nor a success response, e.g. status code < 200 or
// code >= 300 && code <= 399, just dump the raw content into error to help troubleshoot.

type apiNameType int

const apiNameKey apiNameType = iota

// SetTracer set the tracer of opentelemetry.
func (c *Client) SetTracer(tracer trace.Tracer) { _ = "STUB: not implemented"; return }

func withAPIName(ctx context.Context, name string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type UserProfile struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

// GetUserProfile returns the user profile for the specified user.
// Github API doc: https://docs.github.com/en/rest/users/users#get-a-user
func (c *Client) GetUserProfile(ctx context.Context, username string) (user *UserProfile, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Repo struct {
	Name string `json:"name"`
	Star int    `json:"stargazers_count"`
}

// ListUserRepo returns a list of public repositories for the specified user
// Github API doc: https://docs.github.com/en/rest/repos/repos#list-repositories-for-a-user
func (c *Client) ListUserRepo(ctx context.Context, username string, page int) (repos []*Repo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoginWithToken login with GitHub personal access token.
// GitHub API doc: https://docs.github.com/en/rest/overview/other-authentication-methods#authenticating-for-saml-sso
func (c *Client) LoginWithToken(token string) *Client { _ = "STUB: not implemented"; return nil }

// SetDebug enable debug if set to true, disable debug if set to false.
func (c *Client) SetDebug(enable bool) *Client { _ = "STUB: not implemented"; return nil }
