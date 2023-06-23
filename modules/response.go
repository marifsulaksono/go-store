package modules

import (
	"encoding/json"
	"net/http"
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
