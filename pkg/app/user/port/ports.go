package port

import (
	"Project/pkg/app/user/model"
)

type UserServiceInterface interface {
	CreateUserService(u model.CreateUserModel) (succes bool, err error)
}

type UserStorageInterface interface {
	Create(u model.CreateUserModel) (bool, error)
}


