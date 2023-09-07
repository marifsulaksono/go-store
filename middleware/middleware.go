package middleware

import (
	"context"
	"fmt"
	middleError "gostore/helper/domain/errorModel"
	"gostore/helper/response"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id       int
	Username string
	Role     string
	jwt.RegisteredClaims
}

var JWT_SECRET_KEY = []byte("be excited about learning something useful")

// const USERNAME = "arif"
// const PASSWORD = "arf123"

type Key string

const (
	GOSTORE_USERID   Key = "go-store-id"
	GOSTORE_USERNAME Key = "go-store-user"
	GOSTORE_USERROLE Key = "go-store-role"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.URL.Path == "/login" || r.URL.Path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
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
				response.BuildErorResponse(w, middleError.ErrExpToken)
				return
			case jwt.ValidationErrorSignatureInvalid:
				response.BuildErorResponse(w, middleError.ErrUnauthorized)
				return
			default:
				response.BuildErorResponse(w, middleError.ErrUnauthorized)
				return
			}
		}

		if !token.Valid {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		payload, ok := token.Claims.(*JWTClaim)
		if !ok {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		ctx = context.WithValue(ctx, GOSTORE_USERID, payload.Id)
		ctx = context.WithValue(ctx, GOSTORE_USERNAME, payload.Username)
		ctx = context.WithValue(ctx, GOSTORE_USERROLE, payload.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
