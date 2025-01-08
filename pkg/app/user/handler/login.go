package handler

import (
	"Project/internal/api"
	"Project/pkg/app/user/domain"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/validator.v2"
)

func (h UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.LoginUserModel

	body := r.Body
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		api.Error(err, http.StatusBadRequest).Send(w)
		return
	}
	if err := validator.Validate(user); err != nil {
		api.Error(err, http.StatusBadRequest).Send(w)
		return
	}
	token, err := h.UserService.LoginUserService(user)
	if err != nil {
		log.Println(err)

		api.Error(err, http.StatusInternalServerError).Send(w)
		return
	}

	api.Success(http.StatusOK, token).Send(w)
}
