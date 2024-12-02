package handler

import (
	"Project/internal/api"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s PostHandler) GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Fatal(err.Error())
	}

	post, err := s.PostService.GetPostByIdService(id)
	if err != nil {
		api.Error(err, http.StatusInternalServerError)
	}

	api.Succes(http.StatusOK, post).Send(w)
}
