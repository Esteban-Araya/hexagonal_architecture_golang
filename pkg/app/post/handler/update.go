package handler

import (
	"Project/internal/api"
	appError "Project/internal/error"
	"Project/pkg/app/post/domain"
	userDomain "Project/pkg/app/user/domain"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gopkg.in/validator.v2"
)

func (h PostHandler) PostUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var postUpdate domain.PostUpdate

	body := r.Body
	defer body.Close()

	err := json.NewDecoder(body).Decode(&postUpdate)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := validator.Validate(postUpdate); err != nil {
		api.Error(err, http.StatusBadRequest).Send(w)
		return
	}
	user_any := r.Context().Value("user")
	user := user_any.(*userDomain.User)

	err = h.PostService.PostUpdateService(postUpdate, user.ID)

	if err != nil {
		app_error := appError.AppError{}
		if errors.As(err, &app_error) {
			switch app_error {
			case domain.YouHaveNotAccessError:
				api.Error(err, http.StatusUnauthorized).Send(w)
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
