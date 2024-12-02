package post

import (
	"Project/internal/jwt"
	JwtStorage "Project/internal/jwt/storage/postgres"

	Handler "Project/pkg/app/post/handler"
	Service "Project/pkg/app/post/service"
	Storage "Project/pkg/app/post/storage/postgres"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi"
)

func PostRout(db *sql.DB, rout *chi.Mux) {
	jwtStorages := JwtStorage.JwtStorage{DB: db}

	jwtService := jwt.JwtService{jwtStorages}

	JwtMiddleware := jwt.JwtMiddleware{jwtService}

	postStorage := Storage.PostStorage{DB: db}

	postService := Service.PostService{postStorage}

	postHandler := Handler.PostHandler{postService}

	rout.Post("/post", JwtMiddleware.MiddlewareIsJwtValid(http.HandlerFunc(postHandler.CreatePostHandler)))
	rout.Put("/post", JwtMiddleware.MiddlewareIsJwtValid(http.HandlerFunc(postHandler.PostUpdateHandler)))
	rout.Get("/post/{id}", JwtMiddleware.MiddlewareIsJwtValid(http.HandlerFunc(postHandler.GetPostByIdHandler)))
}
