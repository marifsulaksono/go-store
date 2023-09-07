package errorModel

import (
	"gostore/helper/response"
	"net/http"
)

var (
	ErrAddressNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "ADDRESS_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrAddressInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_ADDRESS_INPUT",
		Message:   "Some required input is missing",
	}
	ErrInvalidUser = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_USER",
		Message:   "Unauthorized, invalid user",
	}
	ErrInvalidSA = response.MyErr{
		Status:    http.StatusUnauthorized,
		ErrorCode: "INVALID_SHIPPING_ADDRESS",
		Message:   "Sorry, your shipping address isn't not valid user",
	}
)
