package modules

import (
	"encoding/json"
	"gostore/config"
	"gostore/entity"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

		response := Response{Payload: userRegister, Message: "Register success!"}
		jsonString, errJ := json.Marshal(response)
		if errJ != nil {
			http.Error(w, errJ.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonString)
		return
	}
	http.Error(w, "Method isn't valid!", http.StatusBadRequest)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Unauthorized!"))
	}

	var userLogin entity.User
	if err := config.DB.Where("username = ?", username).First(&userLogin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Username/password is wrong!", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		http.Error(w, "Username/password is wrong!", http.StatusNotFound)
		return
	}

	jwtExpTime := time.Now().Add(time.Hour * 1)
	claims := &JWTClaim{
		Username: userLogin.Username,
		Role:     userLogin.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-store",
			ExpiresAt: jwt.NewNumericDate(jwtExpTime),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenAlgorithm.SignedString(JWT_SECRET_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Data = entity.UserLogin{
		"Name": userLogin.Name,
		"Role": userLogin.Role,
	}

	response := Response{Payload: tokenString, Message: "Login Success!"}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
