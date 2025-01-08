package rout

import (
	"Project/pkg/app/post"
	"Project/pkg/app/user"
	"Project/internal/jwt"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

func StartListener(s *http.Server, db *sql.DB, r *chi.Mux) error {
	jwt.JwtRout(db, r)
	user.UserRout(db, r)
	post.PostRout(db, r)
	return s.ListenAndServe()
}
