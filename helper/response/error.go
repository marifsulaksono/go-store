package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrorInternalServer = MyErr{
		Status:    http.StatusInternalServerError,
		ErrorCode: "INTERNAL_SERVER_ERROR",
		Message:   "Please contact service",
	}
	ErrInvalidParamId = MyErr{
		Status:    http.StatusBadRequest,
		ErrorCode: "INVALID_PARAM_ID",
		Message:   "parameter id isn't valid!",
	}
)

type MyErr struct {
	Status    int            `json:"status_code,omitempty"`
	ErrorCode string         `json:"error_code,omitempty"`
	Message   string         `json:"message,omitempty"`
	Details   map[string]any `json:"detail_error,omitempty"`
}

func (e MyErr) Error() string {
	if e.ErrorCode == "" {
		return e.Message
	}

	if len(e.Details) > 0 {
		return fmt.Sprintf("%s : %v", e.ErrorCode, e.Details)
	}

	return fmt.Sprintf("%s : %v", e.ErrorCode, e.Message)
}

func (e MyErr) Return() error {
	if e.ErrorCode != "" {
		return e
	}

	return nil
}

func (e MyErr) AttachDetail(detail map[string]any) MyErr {
	e.Details = detail
	return e
}

func BuildErorResponse(w http.ResponseWriter, err error) {
	response := buildError(err)
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

func buildError(err error) MyErr {
	var response MyErr
	if checkErr, ok := err.(MyErr); ok {
		response.Status = checkErr.Status
		response.ErrorCode = checkErr.ErrorCode
		response.Message = checkErr.Message
		response.Details = checkErr.Details
	} else {
		response = ErrorInternalServer
	}

	return response
}
