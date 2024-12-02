package service

import (
	"Project/pkg/app/user/domain"
)

func (s UserService) UserUpdateNameService(u domain.UserUpdateName, id_user int) error {

	err := s.UserStorage.UserUpdateName(u, id_user)

	if err != nil {
		return err
	}

	return nil
}
