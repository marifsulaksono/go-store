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

type Response struct {
	Data    interface{}
	Message string
}

func ResponseWrite(w http.ResponseWriter, data interface{}, message string) {
	var response = Response{Data: data, Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func IdVarsMux(w http.ResponseWriter, r *http.Request) (int64, bool) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)
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
		http.Error(w, "Id not found!", http.StatusOK)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
