package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/service"
	"gostore/utils/helper"
	"gostore/utils/helper/domain"
	userError "gostore/utils/helper/domain/errorModel"
	"gostore/utils/response"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{
		Service: s,
	}
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user entity.User
	)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	if err := u.Service.CreateUser(ctx, &user); err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Register success")
}

func (u *UserController) LoginAuth(w http.ResponseWriter, r *http.Request) {
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
		// Basic Authentication
		username, password, ok := r.BasicAuth()
		if !ok {
			response.BuildErorResponse(w, userError.ErrLogin)
			fmt.Println("BasicAuth required")
			return
		}

		// Check user validation on database
		var userLogin entity.UserResponse
		userLogin, err := u.Service.GetUser(ctx, 0, username)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.BuildErorResponse(w, userError.ErrLogin)
				return
			} else {
				response.BuildErorResponse(w, err)
				return
			}
		}

		// Check password validation
		if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
			response.BuildErorResponse(w, userError.ErrLogin)
			return
		}

		token, err := helper.GenerateToken(userLogin)
		if err != nil {
			response.BuildErorResponse(w, err)
			return
		}

		metadata := response.UserInfo{
			Username: userLogin.Username,
			Name:     userLogin.Name,
			Role:     userLogin.Role,
		}

		response.BuildSuccesResponse(w, token, metadata, "Login success")
	}
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
	}

	result, err := u.Service.GetUser(ctx, id, "")
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, result, nil, "")
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user entity.User
	)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	err = u.Service.UpdateUser(ctx, &user)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success update user profile")
}

func (u *UserController) ChangePasswordUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userChange entity.UserChangePassword
	err := json.NewDecoder(r.Body).Decode(&userChange)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}
	defer r.Body.Close()

	err = u.Service.ChangePasswordUser(ctx, userChange)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, nil, nil, "Success change password")
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := u.Service.DeleteUser(ctx)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	message := fmt.Sprintf("Success delete user %d", ctx.Value(helper.GOSTORE_USERID))
	response.BuildSuccesResponse(w, nil, nil, message)
}
