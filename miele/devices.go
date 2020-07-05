package miele

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListDevicesRequest struct {
	// The supported languages for localized values. If the language parameter
	// is missing or invalid, you will receive the english localization.
	// Available values : de, en
	Language string `url:"language,omitempty"`
}

type ListDevicesResponse []Device

func (c *Client) ListDevices(request ListDevicesRequest) (ListDevicesResponse, error) {
	u, err := addOptions("/devices", request)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return ListDevicesResponse{}, err
	}
	var response ListDevicesResponse
	_, err = c.do(req, &response)
	return response, err
}

type DeviceActionRequest struct {
	ProcessAction int  `json:"processAction"`
	Light         int  `json:"light"`
	PowerOn       bool `json:"powerOn"`
	PowerOff      bool `json:"powerOff"`
}

func (c *Client) DeviceAction(deviceID string, request DeviceActionRequest) error {
	u := fmt.Sprintf("/devices/%s/actions", deviceID)
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req, err := c.NewRequest("PUT", u, body)
	if err != nil {
		return err
	}
	resp, err := c.do(req, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
