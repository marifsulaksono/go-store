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
	if r.Method == "POST" {
		var user entity.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashPwd)

		if err := u.Service.CreateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user.Password = ""
		helper.ResponseWrite(w, user, "Register success!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var err error
		// Authentication
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Unauthorized!"))
		}

		// Check user authentication on database
		var userLogin entity.UserResponse
		userLogin, err = u.Service.GetUserByUsername(username)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Username/password is wrong!", http.StatusNotFound)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Check password validation
		if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
			http.Error(w, "Username/password is wrong!", http.StatusNotFound)
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
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func (u *UserController) GetShippingAddressByUserId(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	result, err := u.Service.GetShippingAddressByUserId(userId)
	if err != nil {
		helper.RecordNotFound(w, err)
		return
	}

	message := fmt.Sprintf("Success get shipping address on user %d", userId)
	helper.ResponseWrite(w, result, message)
}

func (u *UserController) InsertShippingAddress(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)

	var SA entity.ShippingAddress
	err := json.NewDecoder(r.Body).Decode(&SA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	SA.UserId = userId
	if err := u.Service.InsertShippingAddress(&SA); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := fmt.Sprintf("Success add shipping address to user %d", userId)
	helper.ResponseWrite(w, SA, message)
}

func (u *UserController) UpdateShippingAddress(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, s := helper.IdVarsMux(w, r); s {
		var SA entity.ShippingAddress
		err := json.NewDecoder(r.Body).Decode(&SA)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		result, err := u.Service.UpdateShippingAddress(userId, id, &SA)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success update shipping address on user %d", userId)
		helper.ResponseWrite(w, result, message)
	}
}

func (u *UserController) DeleteShippingAddress(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.GOSTORE_USERID).(int)
	if id, s := helper.IdVarsMux(w, r); s {
		err := u.Service.DeleteShippingAddress(userId, id)
		if err != nil {
			helper.RecordNotFound(w, err)
			return
		}

		message := fmt.Sprintf("Success delete shipping address on user %d", userId)
		helper.ResponseWrite(w, id, message)
	}
}
