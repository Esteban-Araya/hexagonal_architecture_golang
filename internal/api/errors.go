package api

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Status  int   `json:"-"`
	Message error `json:"error,omitempty"`
}

func Error(err error, status int) *ApiError {
	return &ApiError{
		Status:  status,
		Message: err,
	}
}

func (a ApiError) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(a.Status)
	return json.NewEncoder(w).Encode(a)
}
