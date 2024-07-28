package service

import (
	"Project/pkg/app/user/model"
	"time"
)


func (s UserService) CreateUserService(u model.CreateUserModel) (succes bool, err error){
	u.CreatedAt = time.Now().UTC()

	succes, err = s.Storage.Create(u) 

	return succes, err
}