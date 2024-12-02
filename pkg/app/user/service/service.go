package service

import (
	"Project/pkg/app/user/port"
)

type UserService struct {
	UserStorage port.UserStorageInterface
}
