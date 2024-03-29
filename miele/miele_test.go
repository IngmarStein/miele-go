package miele

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// newTestServer returns a *httptest.Server serving mock responses for the Miele API.
func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/short/devices", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
[
  {
    "fabNumber": "000100000000",
    "state": "Off",
    "type": "Combi Steam Oven",
    "deviceName": "",
    "details": "https://api.mcs3.miele.com/v1/devices/000100000000"
  },
  {
    "fabNumber": "000100000001",
    "state": "Off",
    "type": "Dishwasher",
    "deviceName": "",
    "details": "https://api.mcs3.miele.com/v1/devices/000100000001"
  }
]`))
	})
	mux.HandleFunc("/devices", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
{
  "000100000001": {
    "ident": {
      "type": {
        "key_localized": "Devicetype",
        "value_raw": 7,
        "value_localized": "Dishwasher"
      },
      "deviceName": "",
      "deviceIdentLabel": {
        "fabNumber": "000100000001",
        "fabIndex": "64",
        "techType": "G7965",
        "matNumber": "10992660",
        "swids": [
          "4923",
          "20467",
          "20497",
          "25199",
          "20502",
          "35538",
          "4559",
          "4558",
          "4928",
          "20475",
          "25266",
          "4708",
          "25272",
          "20444",
          "4875",
          "20366",
          "20462"
        ]
      },
      "xkmIdentLabel": {
        "techType": "EK037",
        "releaseVersion": "03.65"
      }
    },
    "state": {
      "ProgramID": {
        "value_raw": 2,
        "value_localized": "Clean Machine",
        "key_localized": "Program Id"
      },
      "status": {
        "value_raw": 2,
        "value_localized": "on",
        "key_localized": "State"
      },
      "programType": {
        "value_raw": 0,
        "value_localized": "Operation mode",
        "key_localized": "Program type"
      },
      "programPhase": {
        "value_raw": 1792,
        "value_localized": "",
        "key_localized": "Phase"
      },
      "remainingTime": [
        0,
        1
      ],
      "startTime": [
        0,
        0
      ],
      "targetTemperature": [
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        },
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        },
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        }
      ],
      "temperature": [
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        },
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        },
        {
          "value_raw": -32768,
          "value_localized": null,
          "unit": "Celsius"
        }
      ],
      "signalInfo": false,
      "signalFailure": false,
      "signalDoor": false,
      "remoteEnable": {
        "fullRemoteControl": true,
        "smartGrid": false
      },
      "light": 2,
      "elapsedTime": [
        0,
        0
      ],
      "spinningSpeed": {
        "unit": "rpm",
        "value_raw": null,
        "value_localized": null,
        "key_localized": "Spinning Speed"
      },
      "dryingStep": {
        "value_raw": null,
        "value_localized": "",
        "key_localized": "Drying level"
      },
      "ventilationStep": {
        "value_raw": null,
        "value_localized": "",
        "key_localized": "Power Level"
      },
      "plateStep": [],
      "ecoFeedback": {
        "currentWaterConsumption": {
          "unit": "l",
          "value": 26.0
        },
		"currentEnergyConsumption": {
          "unit": "kWh",
          "value": 2.0
        },
        "waterForecast": 0.2,
        "energyForecast": 0.1
      },
      "batteryLevel": null
	}
  }
}`))
	})
	return httptest.NewServer(mux)
}

func newTestClient(t *testing.T, svr *httptest.Server) *Client {
	t.Helper()
	baseURL, err := url.Parse(svr.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	return &Client{client: http.DefaultClient, BaseURL: baseURL, UserAgent: userAgent}
}

func TestListShortDevices(t *testing.T) {
	svr := newTestServer()
	defer svr.Close()
	client := newTestClient(t, svr)
	resp, err := client.ListShortDevices(ListShortDevicesRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) != 2 {
		t.Fatalf("unexpected number of devices, expected 2, but got %d", len(resp))
	}

	if resp[0].FabNumber != "000100000000" {
		t.Fatalf("unexpected fab number, expected 000100000000, but got %s", resp[0].FabNumber)
	}
}

func TestListDevices(t *testing.T) {
	svr := newTestServer()
	defer svr.Close()
	client := newTestClient(t, svr)
	resp, err := client.ListDevices(ListDevicesRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) != 1 {
		t.Fatalf("unexpected number of devices, expected 2, but got %d", len(resp))
	}

	d, ok := resp["000100000001"]
	if !ok {
		t.Fatal("did not find device with id 000100000001")
	}
	if d.Ident.Typ.ValueRaw != DEVICE_TYPE_DISHWASHER {
		t.Fatalf("unexpected device type, expected %d, but got %d", DEVICE_TYPE_DISHWASHER, d.Ident.Typ.ValueRaw)
	}
}
