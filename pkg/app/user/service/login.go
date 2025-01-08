package service

import (
	"Project/internal/encryption"
	"Project/internal/jwt"
	"Project/pkg/app/user/domain"
)	

func (s UserService) LoginUserService(u domain.LoginUserModel) (*jwt.RefreshAndAccessTokens, error) {

	user, err := s.UserStorage.GetUserByEmail(u.Email)
	if err != nil {

		return nil, err
	}

	decrypted_password, err := encryption.Decrypt(user.Password)
	if err != nil {
		return nil, err
	}

	if string(decrypted_password) != u.Password {
		return nil, domain.EmailOrPasswordIsWrong
	}

	access, err := jwt.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refresh, err := jwt.GenerateRefreshToken(user)

	if err != nil {
		return nil, err
	}
	tokens := jwt.RefreshAndAccessTokens{
		Refresh: refresh,
		Access: access,
	}

	if err != nil {
		return nil, err
	}

	return &tokens, nil
}
