package service

import (
	"Project/pkg/app/user/port"
)
type UserService struct{
	Storage port.UserStorageInterface
}