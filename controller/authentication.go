package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/service"
	"gostore/utils/helper/domain"
	"gostore/utils/response"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type AuthenticationController struct {
	Service service.AuthenticationService
}

func NewAuthenticationController(s service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{Service: s}
}

func (a *AuthenticationController) LoginController(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	auth := mux.Vars(r)["auth"]
	if auth == "google" {
		// Google OAuth
		URL, err := url.Parse(domain.OAuthGoogleConf.Endpoint.AuthURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set required parameters
		parameters := url.Values{}
		parameters.Add("client_id", domain.OAuthGoogleConf.ClientID)
		parameters.Add("scope", strings.Join(domain.OAuthGoogleConf.Scopes, " "))
		parameters.Add("redirect_uri", domain.OAuthGoogleConf.RedirectURL)
		parameters.Add("response_type", "code")
		parameters.Add("state", domain.OAuthStateString)
		URL.RawQuery = parameters.Encode()
		url := URL.String()
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else if auth == "" {
		var credential entity.Credential
		if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
			fmt.Println(err)
			response.BuildErorResponse(w, err)
			return
		}

		accessToken, refreshToken, err := a.Service.LoginService(ctx, &credential)
		if err != nil {
			response.BuildErorResponse(w, err)
			return
		}

		data := response.LoginInfo{
			Username:     credential.Username,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		response.BuildSuccesResponse(w, data, nil, "Login success")
	}
}

func (a *AuthenticationController) RenewAccessToken(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		s   entity.Session
	)

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	accessToken, err := a.Service.RenewAccessToken(ctx, &s)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	data := response.LoginInfo{
		AccessToken: accessToken,
	}

	response.BuildSuccesResponse(w, data, nil, "Success create new access token")
}

func (a *AuthenticationController) LogoutController(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		s   entity.Session
	)

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	if err := a.Service.LogoutService(ctx, &s); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Logout success..")
}

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
