package config

import (
	"gostore/utils/helper/domain"
	"os"
)

func InitGoogleConfig() {
	domain.OAuthGoogleConf.ClientID = os.Getenv("CLIENT_ID")
	domain.OAuthGoogleConf.ClientSecret = os.Getenv("CLIENT_SECRET")
	domain.OAuthGoogleConf.RedirectURL = os.Getenv("REDIRECT_URL")
	domain.OAuthStateString = os.Getenv("STATE_STRING")
}
