package controller

import (
	"encoding/json"
	"fmt"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"gostore/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := u.Service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = ""
	helper.ResponseWrite(w, user, "Register success!")
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	// Basic Authentication
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Unauthorized!"))
	}

	// Check user validation on database
	var userLogin entity.UserResponse
	fmt.Println(username)
	userLogin, err = u.Service.GetUser(0, username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Username/password is wrong!", http.StatusUnauthorized)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Check password validation
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		http.Error(w, "Username/password is wrong!", http.StatusUnauthorized)
		return
	}

	// Create token claim
	jwtExpTime := time.Now().Add(time.Hour * 24)
	claims := &middleware.JWTClaim{
		Id:       userLogin.Id,
		Username: userLogin.Username,
		Role:     userLogin.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-store",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
		},
	}

	// Generate JWT Token
	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenAlgorithm.SignedString(middleware.JWT_SECRET_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Login success, welcome %s!", userLogin.Name)
	helper.ResponseWrite(w, tokenString, message)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.Service.UpdateUser(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, user, "Success update user profile")
}

func (u *UserController) ChangePasswordUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userChange entity.UserChangePassword
	err := json.NewDecoder(r.Body).Decode(&userChange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.Service.ChangePasswordUser(ctx, userChange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, nil, "Success change password")
}

func (u *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := u.Service.DeleteUser(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success delete user %d", ctx.Value(middleware.GOSTORE_USERID))
	helper.ResponseWrite(w, nil, message)
}

func (u *UserController) GetShippingAddressByUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := u.Service.GetShippingAddressByUserId(ctx)
	if err != nil {
		helper.RecordNotFound(w, err)
		return
	}

	helper.ResponseWrite(w, result, "Success get shipping address")
}

func (u *UserController) InsertShippingAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var SA entity.ShippingAddress
	err := json.NewDecoder(r.Body).Decode(&SA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := u.Service.InsertShippingAddress(ctx, &SA); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.ResponseWrite(w, SA, "Success add shipping address")
}

func (u *UserController) UpdateShippingAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		var SA entity.ShippingAddress
		err := json.NewDecoder(r.Body).Decode(&SA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		result, err := u.Service.UpdateShippingAddress(ctx, id, &SA)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		helper.ResponseWrite(w, result, "Success update shipping address")
	}
}

func (u *UserController) DeleteShippingAddress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if id, s := helper.IdVarsMux(w, r); s {
		err := u.Service.DeleteShippingAddress(ctx, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		helper.ResponseWrite(w, id, "Success delete shipping address")
	}
}
