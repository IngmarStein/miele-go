package miele

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocalizedRequest struct {
	// The supported languages for localized values. If the language parameter
	// is missing or invalid, you will receive the english localization.
	// Available values : de, en
	Language string `url:"language,omitempty"`
}

type ListDevicesRequest struct {
	LocalizedRequest
}

type ListDevicesResponse map[string]Device

func (c *Client) ListDevices(request ListDevicesRequest) (ListDevicesResponse, error) {
	u, err := addOptions("devices", request)
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

type GetDeviceRequest struct {
	LocalizedRequest
}

func (c *Client) GetDevice(deviceID string, request GetDeviceRequest) (Device, error) {
	u, err := addOptions(fmt.Sprintf("devices/%s/state", deviceID), request)
	if err != nil {
		return Device{}, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return Device{}, err
	}
	var response Device
	_, err = c.do(req, &response)
	return response, err
}

type GetDeviceStateRequest struct {
	LocalizedRequest
}

func (c *Client) GetDeviceState(deviceID string, request GetDeviceStateRequest) (State, error) {
	u, err := addOptions(fmt.Sprintf("devices/%s/state", deviceID), request)
	if err != nil {
		return State{}, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return State{}, err
	}
	var response State
	_, err = c.do(req, &response)
	return response, err
}

type GetDeviceIdentRequest struct {
	LocalizedRequest
}

func (c *Client) GetDeviceIdent(deviceID string, request GetDeviceStateRequest) (Ident, error) {
	u, err := addOptions(fmt.Sprintf("devices/%s/ident", deviceID), request)
	if err != nil {
		return Ident{}, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return Ident{}, err
	}
	var response Ident
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
	u := fmt.Sprintf("devices/%s/actions", deviceID)
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
