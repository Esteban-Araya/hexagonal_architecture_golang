package jwt

import (
	userDomain "Project/pkg/app/user/domain"

	jwt_lib "github.com/golang-jwt/jwt/v5"
)

type JwtServiceInterface interface {
	ValidateRefreshToken(value string) (jwt_lib.MapClaims, error)
	ValidateAccessToken(value string) (jwt_lib.MapClaims, error)	
	IsTokenExpire(time_token float64) error
	GetUser(id float64) (*userDomain.User, error)
}

type JwtStorageInterface interface {
	GetUserById(id float64) (*userDomain.User, error)
}
