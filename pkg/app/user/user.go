package user

import (
	"Project/internal/jwt"
	JwtStorage "Project/internal/jwt/storage/postgres"

	Handler "Project/pkg/app/user/handler"
	Service "Project/pkg/app/user/service"
	Storage "Project/pkg/app/user/storage/postgres"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi"
)

func UserRout(db *sql.DB, rout *chi.Mux) {
	jwtStorages := JwtStorage.JwtStorage{DB: db}

	jwtService := jwt.JwtService{JwtStorage: jwtStorages}

	JwtMiddleware := jwt.JwtMiddleware{JwtService: jwtService}

	userStorage := Storage.UserStorage{DB: db}

	userService := Service.UserService{UserStorage: userStorage}

	userHandler := Handler.UserHandler{UserService: userService}

	rout.Post("/user", userHandler.CreateUserHandler)
	rout.Post("/user/login", userHandler.LoginUserHandler)
	rout.Get("/user", JwtMiddleware.MiddlewareIsAccessTokenValid(http.HandlerFunc(userHandler.GetUser)))
	rout.Put("/user", JwtMiddleware.MiddlewareIsAccessTokenValid(http.HandlerFunc(userHandler.UserUpdateHandler)))
}
