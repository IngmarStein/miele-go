package miele

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.mcs3.miele.com/v1/"
	userAgent      = "go-miele"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	client *http.Client
}

// NewClient returns a new Miele API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Token: token}

	return c
}
