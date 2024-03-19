package controller

import (
	"encoding/json"
	"gostore/utils/helper/domain"
	"gostore/utils/response"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

func CallbackGoogleAuth(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	log.Println("New Callback from Google OAuth :\n" + state)
	if state != domain.OAuthStateString {
		log.Printf("Invalid state. expected %s, got %s\n", domain.OAuthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		log.Println("Code not found")
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User permission has denied"))
			return
		}

		w.Write([]byte("Code Not Found to provide AccessToken"))
	} else {
		token, err := domain.OAuthGoogleConf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Printf("OAuth Exchange failed : %v\n", err)
			return
		}

		log.Printf("[TOKEN_AUTH]Access Token : %s", token.AccessToken)
		log.Printf("[TOKEN_AUTH]Expiry Token : %s", token.Expiry.String())
		log.Printf("[TOKEN_AUTH]Refresh Token : %s", token.RefreshToken)

		responseBody, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			log.Printf("Error Get Response : %v", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer responseBody.Body.Close()

		var responseJSON map[string]any
		if err := json.NewDecoder(responseBody.Body).Decode(&responseJSON); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.BuildSuccesResponse(w, responseJSON, nil, "")
		return
	}
}
