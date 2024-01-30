package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
	ErrCartNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "CART_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrCantAddToCart = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_ADD_CART",
		Message:   "cannot add your product's store to your cart",
	}
	ErrStockProductNotEnough = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "STOCK_NOT_ENOUGH",
		Message:   "stock product isn't enough",
	}
	ErrCartInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_CART_INPUT",
		Message:   "Some required input is missing",
	}
)
