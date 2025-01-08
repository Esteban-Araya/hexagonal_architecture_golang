package jwt

import (
	
	"Project/internal/api"
	"net/http"
	"Project/pkg/app/user/domain"

)


type JwtHandler struct {
}

func (h JwtHandler) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	
	user := r.Context().Value("user").(*domain.User)

	access, err := GenerateAccessToken(user)
	if err != nil {

		api.Error(err, http.StatusInternalServerError).Send(w)
		return
	}

	api.Success(http.StatusOK, access).Send(w)
}
