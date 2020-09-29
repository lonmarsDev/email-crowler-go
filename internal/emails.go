package internal

import (
	"fmt"

	"github.com/realpamisa/gmail-extractor/api"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/endpoint",
	ClientID:     "409551055881-09tt2lprbfhuiurd7jr5q74p184ougp4.apps.googleusercontent.com",
	ClientSecret: "4KRHQrXW-ETZRh5eGDXPAwAs",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://mail.google.com/"},
	Endpoint:     google.Endpoint,
}

var OauthStateString = "state-token"

func ExtractEmail(state string, code string) ([]string, error) {
	if state != OauthStateString {
		fmt.Errorf("invalid oauth state")
		return nil, nil
	}

	//get google access token
	token, err := GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Errorf("code exchange failed: %s", err.Error())
		return nil, err
	}

	//get userid
	userID, err := GetUserID(token.AccessToken)
	if err != nil {
		panic(err)

		return nil, err
	}
	//get all emails
	api.GetAllMails(token.AccessToken, userID)
	if err != nil {
		panic(err)
		return nil, err
	}

	return nil, err

}
