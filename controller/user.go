package controller

import (
	"encoding/json"
	"fmt"
	"gostore/config"
	"gostore/entity"
	"gostore/helper"
	"gostore/middleware"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var userRegister entity.User
		err := json.NewDecoder(r.Body).Decode(&userRegister)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		hashPwd, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
		userRegister.Password = string(hashPwd)

		if err := config.DB.Create(&userRegister).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helper.ResponseWrite(w, userRegister, "Register success!")
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var err error
		// Authentication
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Unauthorized!"))
		}

		// Check user authentication on database
		var userLogin entity.User
		userLogin, err = entity.GetUserValid(username)
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
		jwtExpTime := time.Now().Add(time.Hour * 1)
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
