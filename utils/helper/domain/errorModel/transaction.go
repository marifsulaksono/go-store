package errorModel

import (
	"gostore/utils/response"
	"net/http"
)

var (
	ErrTransactionNotFound = response.MyErr{
		Status:    http.StatusNotFound,
		ErrorCode: "TRANSACTION_NOT_FOUND",
		Message:   "The ID provided is not found",
	}
	ErrTransactionInput = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_TRANSACTION_INPUT",
		Message:   "Some required input is missing",
	}
	ErrCantAddToTrx = response.MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_ADD_TRANSACTION",
		Message:   "cannot add your product's store to your transaction",
	}
)
