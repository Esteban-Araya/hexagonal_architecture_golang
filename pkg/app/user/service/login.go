package service

import (
	"Project/internal/encryption"
	"Project/internal/jwt"
	"Project/pkg/app/user/domain"
)

func (s UserService) LoginUserService(u domain.LoginUserModel) (string, error) {

	user, err := s.UserStorage.GetUserByEmail(u.Email)
	if err != nil {

		return "", err
	}

	decrypted_password, err := encryption.Decrypt(user.Password)
	if err != nil {
		return "", err
	}

	if string(decrypted_password) != u.Password {
		return "", domain.EmailOrPasswordIsWrong
	}

	token, err := jwt.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
