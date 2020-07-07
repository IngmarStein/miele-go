package miele

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.mcs3.miele.com/thirdparty/login",
	TokenURL: "https://api.mcs3.miele.com/thirdparty/token",
}
