package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
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
