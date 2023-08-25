package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type response struct {
	Data     interface{}   `json:"data,omitempty"`
	Metadata interface{}   `json:"metadata,omitempty"`
	Message  string        `json:"message,omitempty"`
	Code     string        `json:"error_code,omitempty"`
	Details  []DetailError `json:"detail_error,omitempty"`
}

type DetailError struct {
	Field string `json:"field,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func buildResponseSuccess(data interface{}, metadata interface{}, message string) response {
	var res response
	res.Data = data
	res.Metadata = metadata
	res.Message = message

	return res
}

func buildResponseError(message string, code string, err []DetailError) response {
	var res response
	res.Code = code
	res.Message = message
	res.Details = err

	return res
}

func BuildResponseSuccess(w http.ResponseWriter, data interface{}, metadata interface{}, message string) {
	payload := buildResponseSuccess(data, metadata, message)
	jsonData, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func BuildError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		buildErrNotFound(w, ProductNotFoundError, "the ID provided is not found")
	case isBadRequestError(err):
		buildBadRequest(w, BadRequestError, err.Error())
	case isUnauthorized(err):
		buildErrUnauthorized(w, UnauthorizedError, err)
	default:
		buildErrInternalServer(w, InternalServerError, err)
	}
}

func buildErrNotFound(w http.ResponseWriter, code string, message string) {
	var response = buildResponseError(message, code, nil)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonData)
}

func buildBadRequest(w http.ResponseWriter, code string, message string) {
	var response = buildResponseError(message, code, nil)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonData)
}

func buildErrUnauthorized(w http.ResponseWriter, code string, err error) {
	var response = buildResponseError(err.Error(), code, nil)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonData)
}

func buildErrInternalServer(w http.ResponseWriter, code string, err error) {
	var response = buildResponseError(err.Error(), code, nil)
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonData)
}

func isBadRequestError(err error) bool {
	var a bool

	switch err {
	case ErrRecDeleted:
		a = true
	case ErrRecRestored:
		a = true
	case ErrUserExist:
		a = true
	case ErrStockNotEnough:
		a = true
	case ErrInvalidSA:
		a = true
	case ErrDuplicateStore:
		a = true
	case ErrAddProductTo:
		a = true
	case ErrWrongOldPassword:
		a = true
	default:
		a = false
	}

	return a
}

func isUnauthorized(err error) bool {
	var b bool
	switch err {
	case ErrAccDeny:
		b = true
	case ErrInvalidUser:
		b = true
	case ErrLoginAcc:
		b = true
	default:
		b = false
	}

	return b
}

func ResponseWrite(w http.ResponseWriter, data interface{}, message string) {
	var response = response{Data: data, Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func IdVarsMux(w http.ResponseWriter, r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("parameter id isn't valid!"))
		fmt.Println(err.Error())
		return 0, false
	}
	return id, true
}

func RecordNotFound(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
