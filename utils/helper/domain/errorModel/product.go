package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
	ErrProductNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "PRODUCT_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrProductInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_PRODUCT_INPUT",
		Message:   "Some required input is missing",
	}
	ErrProductDeleted = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "PRODUCT_DELETED",
		Message:   "The product ID provided was deleted",
	}
	ErrProductRestored = response.MyErr{
		Status:    http.StatusOK,
		ErrorCode: "PRODUCT_RESTORED",
		Message:   "The product ID provided was restored",
	}
)
