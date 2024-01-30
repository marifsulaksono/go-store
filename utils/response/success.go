package response

import (
	"encoding/json"
	"net/http"
)

type MySuccess struct {
	Status   int         `json:"status_code,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Message  string      `json:"message,omitempty"`
}

func BuildSuccesResponse(w http.ResponseWriter, data, metadata interface{}, message string) {
	payload := buildSuccess(data, metadata, message)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func buildSuccess(data, metadata interface{}, message string) MySuccess {
	var response MySuccess
	response.Status = http.StatusOK
	response.Data = data
	response.Metadata = metadata
	response.Message = message

	return response
}
