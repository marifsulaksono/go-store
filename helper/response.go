package helper

import (
	"errors"
	"net/http"
	"strconv"

	responseEror "gostore/helper/response"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ParamIdChecker(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, responseEror.ErrInvalidParamId
	}
	return id, nil
}

func RecordNotFound(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Id not found", http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
