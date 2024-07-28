package user

import (
    "database/sql"
    Storage "Project/pkg/app/user/storage/postgres"
    Service "Project/pkg/app/user/service"
    Handler "Project/pkg/app/user/handler"
    "net/http"

)

func UserRout(db *sql.DB){
    userStorage := Storage.UserStorage{DB: db}

    service := Service.UserService{userStorage}

    handler := Handler.UserHandler{service}

    http.HandleFunc("/user", handler.CreatUserHandler)
}