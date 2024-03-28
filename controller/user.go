package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/service"
	"gostore/utils/helper"
	"gostore/utils/response"
	"net/http"
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

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	id, err := helper.ParamIdChecker(w, r)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
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
