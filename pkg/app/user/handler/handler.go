package user

import (
	"Project/pkg/app/user/port"
)
type UserHandler struct{
	UserService port.UserServiceInterface
}