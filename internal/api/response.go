package api

import (
	"encoding/json"
	"net/http"
)

type ApiMessage struct {
	Status  int         `json:"-"`
	Message interface{} `json:"message,omitempty"`
}

func Success(status int, obj ...interface{}) *ApiMessage {

	if obj == nil {
		return &ApiMessage{
			Status: status,
		}
	}

	return &ApiMessage{
		Status:  status,
		Message: obj[0],
	}
}

func (a ApiMessage) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(a.Status)
	return json.NewEncoder(w).Encode(a.Message)
}
