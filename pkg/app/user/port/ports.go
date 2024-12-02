package port

import (
	"Project/pkg/app/user/domain"
)

type UserServiceInterface interface {
	CreateUserService(u domain.CreateUserModel) (err error)
	LoginUserService(u domain.LoginUserModel) (string, error)
	UserUpdateNameService(u domain.UserUpdateName, id_user int) error
}

type UserStorageInterface interface {
	Create(u domain.CreateUserModel) error
	GetUserByEmail(email string) (*domain.User, error)
	UserUpdateName(u domain.UserUpdateName, id_user int) error
}
