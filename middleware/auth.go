package middleware

import (
	"context"
	"gostore/utils/helper"
	middleError "gostore/utils/helper/domain/errorModel"
	"gostore/utils/response"
	"net/http"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenString, ok := helper.GetTokenFromHeader(w, r)
		if !ok {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		payload, ok, err := helper.ValidateJWT(tokenString)
		if !ok || err != nil {
			response.BuildErorResponse(w, middleError.ErrInvalidToken)
			return
		}

		ctx = context.WithValue(ctx, helper.GOSTORE_USERID, payload.Id)
		ctx = context.WithValue(ctx, helper.GOSTORE_USERNAME, payload.Username)
		ctx = context.WithValue(ctx, helper.GOSTORE_USEREMAIL, payload.Email)
		ctx = context.WithValue(ctx, helper.GOSTORE_USERROLE, payload.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
