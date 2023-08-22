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
	Data    interface{}   `json:"data,omitempty"`
	Message string        `json:"message,omitempty"`
	Code    string        `json:"error_code,omitempty"`
	Details []DetailError `json:"detail_error,omitempty"`
}

type DetailError struct {
	Field string `json:"field,omitempty"`
	Desc  string `json:"desc,omitempty"`
}

func buildResponse(data interface{}, message string, code string, err []DetailError) response {
	var res response
	res.Data = data
	res.Code = code
	res.Message = message
	res.Details = err

	return res
}

func BuildResponseSuccess(w http.ResponseWriter, data interface{}, message string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buildResponse(data, message, "", nil))
}

func BuildError(w http.ResponseWriter, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		buildErrNotFound(w, ProductNotFoundError, "the ID provided is not found")
	default:
		buildErrInternalServer(w, InternalServerError, err)
	}
}

func buildErrNotFound(w http.ResponseWriter, code string, message string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(buildResponse(nil, message, code, nil))
}

func buildErrInternalServer(w http.ResponseWriter, code string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(buildResponse(nil, err.Error(), code, []DetailError{}))
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
