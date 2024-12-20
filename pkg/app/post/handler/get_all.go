package handler

import (
	"Project/internal/api"
	"log"
	"net/http"
)

func (s PostHandler) GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ME HAN LLAMAO TIO")
	post, err := s.PostService.GetAllPostService()
	if err != nil {
		api.Error(err, http.StatusInternalServerError)
	}

	api.Succes(http.StatusOK, post).Send(w)
}
