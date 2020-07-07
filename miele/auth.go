package miele

import (
	"fmt"
	"net/http"
	"net/url"
)

type GetTokenRequest struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	Locale       string
}

func (c *Client) GetToken(request GetTokenRequest) error {
	data := url.Values{}
	data.Set("client_id", request.ClientID)
	data.Set("client_secret", request.ClientSecret)
	data.Set("grant_type", "password")
	data.Set("username", request.Username)
	data.Set("password", request.Password)
	data.Set("state", "token")
	data.Set("redirect_uri", "/v1/devices")
	data.Set("vg", request.Locale)

	req, err := c.NewRequest("POST", "thirdparty/token", data.Encode())
	if err != nil {
		return err
	}
	resp, err := c.do(req, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
