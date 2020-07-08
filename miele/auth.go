package miele

import (
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.mcs3.miele.com/thirdparty/login",
	TokenURL: "https://api.mcs3.miele.com/thirdparty/token",
}

// AuthTransport can be used to add the required custom fields for an
// OAuth2 client using golang.org/x/oauth2.
// VG is the locale used when registering the Miel@Home account.
//
// Example:
// hc := &http.Client{Transport: &miele.AuthTransport{VG: *vg}}
// ctx := context.WithValue(context.Background(), oauth2.HTTPClient, hc)
type AuthTransport struct {
	VG string
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	vals, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	vals.Set("vg", t.VG)

	buf := strings.NewReader(vals.Encode())
	req.Body = ioutil.NopCloser(buf)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
	req.ContentLength = int64(buf.Len())

	// Call default roundtrip
	return http.DefaultTransport.RoundTrip(req)
}
