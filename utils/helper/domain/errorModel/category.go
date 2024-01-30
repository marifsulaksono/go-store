package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
	ErrCategoryNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "CATEGORY_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrCategoryInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_CATEGORY_INPUT",
		Message:   "Some required input is missing",
	}
)
