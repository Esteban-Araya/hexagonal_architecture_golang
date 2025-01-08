package jwt

import (
	"net/http"
	"os"
	"database/sql"
	Storage "Project/internal/jwt/storage/postgres"

	"github.com/go-chi/chi"
)

func JwtRout(db *sql.DB, rout *chi.Mux) {
	ACCESS_SECRET_KEY = os.Getenv("ACCESS_SECRET_KEY")
	REFRESH_SECRET_KEY = os.Getenv("REFRESH_SECRET_KEY")

	jwtStorage := Storage.JwtStorage{DB: db}
	jwtService := JwtService{jwtStorage}
	jwtMiddleware := JwtMiddleware{jwtService}
	jwtHandler := JwtHandler{}
	rout.Get("/acces_token", jwtMiddleware.MiddlewareIsRefreshTokenValid(http.HandlerFunc(jwtHandler.GetAccessToken)))
}

