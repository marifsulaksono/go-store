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
	ErrExistUser = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "USER_EXIST",
		Message:   "Username or email already exist",
	}
	ErrLogin = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_LOGIN",
		Message:   "Username or password is wrong",
	}
	ErrUserNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "USER_NOT_FOUND",
		Message:   "The id or username provided is not found",
	}
	ErrWrongPassword = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_CURRENT_PASS",
		Message:   "Wrong current password",
	}
	ErrUserInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_USER_INPUT",
		Message:   "Some required input is missing",
	}
	ErrUserDeleted = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "USER_DELETED",
		Message:   "The user ID provided was deleted",
	}
	ErrUserRestored = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "USER_RESTORED",
		Message:   "The user ID provided was restored",
	}
)
