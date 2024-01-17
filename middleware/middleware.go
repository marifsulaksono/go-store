package middleware

import (
	"context"
	"gostore/helper"
	middleError "gostore/helper/domain/errorModel"
	"gostore/helper/response"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.HandlerFunc, allowedRole ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.URL.Path == "/login" || r.URL.Path == "/register" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := helper.GetTokenFromHeader(w, r)
		if tokenString == "" {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		token, err := helper.ValidateJWT(tokenString)
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

		payload, ok := token.Claims.(*helper.JWTClaim)
		if !ok {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		if len(allowedRole) > 0 {
			allow := helper.RoleBasedAccessControl(r, allowedRole...)
			if !allow {
				response.BuildErorResponse(w, middleError.ErrNotAllowed)
				return
			}
		}

		ctx = context.WithValue(ctx, helper.GOSTORE_USERID, payload.Id)
		ctx = context.WithValue(ctx, helper.GOSTORE_USERNAME, payload.Username)
		ctx = context.WithValue(ctx, helper.GOSTORE_USERROLE, payload.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
