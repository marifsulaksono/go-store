package middleware

import (
	"gostore/utils/helper"
	middleError "gostore/utils/helper/domain/errorModel"
	"gostore/utils/response"
	"net/http"
)

func RBACAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role := ctx.Value(helper.GOSTORE_USERROLE).(string)

		if role != "admin" {
			response.BuildErorResponse(w, middleError.ErrNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RBACSellerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role := ctx.Value(helper.GOSTORE_USERROLE).(string)

		if role != "seller" {
			response.BuildErorResponse(w, middleError.ErrNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}
