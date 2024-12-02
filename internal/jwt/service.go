package jwt

import (
	"Project/pkg/app/user/domain"
	userDomain "Project/pkg/app/user/domain"
	"os"
	"time"

	jwt_lib "github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	JwtStorage JwtStorageInterface
}

func GenerateToken(user *domain.User) (string, error) {

	claims := JWT{
		user.ID,
		jwt_lib.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt_lib.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s JwtService) ValidateToken(value string) (jwt_lib.MapClaims, error) {
	token, err := jwt_lib.Parse(value, func(t *jwt_lib.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt_lib.MapClaims), nil
}

func (s JwtService) IsTokenExpire(time_token float64) error {

	if time_token < float64(jwt_lib.NewNumericDate(time.Now()).Second()) {
		return TokenExpireError
	}

	return nil
}

func (s JwtService) GetUser(id float64) (*userDomain.User, error) {

	user, err := s.JwtStorage.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, err
}
