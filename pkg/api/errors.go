package api

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Status 	   int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func Error(err error, status int) *ApiError{
	return &ApiError{
		Status: status,
		Type: "error",
		Message: err.Error(),
	}
}

func (a ApiError) Send (w http.ResponseWriter) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(a.Status)
	return json.NewEncoder(w).Encode(a)
}