package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
	ErrStoreNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "STORE_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrDuplicateStore = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "DUPLICATE_STORE",
		Message:   "User allowed has one store only",
	}
	ErrInvalidUserStore = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_USER_STORE",
		Message:   "Invalid user, only store owner can access",
	}
	ErrStoreInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_USER_INPUT",
		Message:   "Some required input is missing",
	}
	ErrStoreDeleted = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "STORE_DELETED",
		Message:   "The Store ID provided was deleted",
	}
	ErrStoreRestored = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "STORE_RESTORED",
		Message:   "The Store ID provided was restored",
	}
)
