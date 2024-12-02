package service

import (
	"Project/internal/encryption"
	"Project/pkg/app/user/domain"

	"time"
)

func (s UserService) CreateUserService(u domain.CreateUserModel) (err error) {

	if u.Password != u.PasswordVerified {
		return domain.DiferentPasswordError
	}
	password_encrypted, err := encryption.Encrypt([]byte(u.Password))

	if err != nil {
		return err
	}

	u.Password = password_encrypted
	u.CreatedAt = time.Now().UTC()

	err = s.UserStorage.Create(u)

	return err
}
