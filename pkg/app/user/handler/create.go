package handler

import (
	"Project/internal/api"
	appError "Project/internal/error"
	"Project/pkg/app/user/domain"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gopkg.in/validator.v2"
)

func (h UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.CreateUserModel

	// response := make(map[string]bool)

	body := r.Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := validator.Validate(user); err != nil {
		api.Error(err, http.StatusBadRequest).Send(w)
		return
	}

	err = h.UserService.CreateUserService(user)

	if err != nil {
		app_error := appError.AppError{}
		if errors.As(err, &app_error) {
			switch app_error {
			case domain.DiferentPasswordError:
				api.Error(err, http.StatusBadRequest).Send(w)
			case domain.EmailAlreadyExistError:
				api.Error(err, http.StatusBadRequest).Send(w)
			case domain.DatabaseError:
				api.Error(domain.ServerError, http.StatusInternalServerError).Send(w)
			default:
				api.Error(domain.ServerError, http.StatusInternalServerError).Send(w)
			}
			return
		}
		api.Error(domain.ServerError, http.StatusInternalServerError).Send(w)
		return
	}

	api.Success(http.StatusCreated).Send(w)
}
