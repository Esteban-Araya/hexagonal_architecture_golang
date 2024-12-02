package handler

import (
	"Project/internal/api"
	"Project/pkg/app/user/domain"
	"encoding/json"
	"log"
	"net/http"
)

func (h UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.LoginUserModel

	body := r.Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		log.Fatal(err.Error())
	}
	token, err := h.UserService.LoginUserService(user)
	if err != nil {
		api.Error(err, http.StatusInternalServerError).Send(w)
	}
	api.Succes(http.StatusOK, token).Send(w)
}
