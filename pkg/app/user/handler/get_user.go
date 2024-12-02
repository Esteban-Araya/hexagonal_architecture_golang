package handler

import (
	"Project/internal/api"

	"net/http"
)

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// var user domain.User
	user := r.Context().Value("user")

	api.Succes(http.StatusAccepted, user).Send(w)
}
