package utils

import (
	"five/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitConfig(config utils.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  config.RedirectUri,
	}
}
