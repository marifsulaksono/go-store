package modules

import (
	"fmt"
	"gostore/entity"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Username string
	Role     string
	jwt.RegisteredClaims
}

// const USERNAME = "arif"
// const PASSWORD = "arf123"

var JWT_SECRET_KEY = []byte("be excited about learning something useful")
var Data entity.UserLogin

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		claims := &JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return JWT_SECRET_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorExpired:
				http.Error(w, "Unauthorized, token expired!", http.StatusUnauthorized)
				return
			case jwt.ValidationErrorSignatureInvalid:
				http.Error(w, "Unauthorized!", http.StatusUnauthorized)
				return
			default:
				http.Error(w, "Unauthorized!", http.StatusUnauthorized)
				return
			}
		}

		if !token.Valid {
			http.Error(w, "Unauthorized!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
