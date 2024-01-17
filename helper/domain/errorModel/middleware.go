package errorModel

import (
	"gostore/helper/response"
	"net/http"
)

var (
	ErrInvalidToken = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_TOKEN",
		Message:   "The token provided isn't valid",
	}
	ErrUnauthorized = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_AUTHORIZE",
		Message:   "Need valid authorization",
	}
	ErrExpToken = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "EXPIRED_TOKEN",
		Message:   "Token expired, please login again",
	}
	ErrNotAllowed = response.MyErr{
		Status:    http.StatusForbidden,
		ErrorCode: "FORBIDDEN_ACCESS",
		Message:   "This endpoint is not allowed",
	}
)
