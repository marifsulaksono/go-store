package domain

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OAuthGoogleConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	OAuthStateString = ""
)
