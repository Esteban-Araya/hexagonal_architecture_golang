package jwt

import (
	"Project/pkg/app/user/domain"
	"os"
	"time"
	"log"
	jwt_lib "github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	JwtStorage JwtStorageInterface
}

var ACCESS_SECRET_KEY = ""
var REFRESH_SECRET_KEY = ""
func GenerateToken(user *domain.User, secret_key string, hours_to_add time.Duration) (string, error) {

	claims := JWT{
		user.ID,
		jwt_lib.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt_lib.NewNumericDate(time.Now().Add(hours_to_add * time.Hour)),
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret_key))
}

func GenerateAccessToken(user *domain.User) (string, error) {
 
	return GenerateToken(user,ACCESS_SECRET_KEY, 20)
}

func GenerateRefreshToken(user *domain.User) (string, error) {
 
	return GenerateToken(user,REFRESH_SECRET_KEY, 120)
}

func ValidateToken(value string, secret_key string) (jwt_lib.MapClaims, error) {
	token, err := jwt_lib.Parse(value, func(t *jwt_lib.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	if err != nil {
		return nil, TokenIsNotValidError
	}

	return token.Claims.(jwt_lib.MapClaims), nil
}

func (s JwtService) ValidateAccessToken(value string) (jwt_lib.MapClaims, error) {
	
	return ValidateToken(value, ACCESS_SECRET_KEY)
}


func (s JwtService) ValidateRefreshToken(value string) (jwt_lib.MapClaims, error) {
	log.Println(REFRESH_SECRET_KEY)
	log.Println(os.Getenv("REFRESH_SECRET_KEY"))

	return ValidateToken(value, REFRESH_SECRET_KEY)
}


func (s JwtService) IsTokenExpire(time_token float64) error {

	if time_token < float64(jwt_lib.NewNumericDate(time.Now()).Second()) {
		return TokenExpireError
	}

	return nil
}

func (s JwtService) GetUser(id float64) (*domain.User, error) {

	user, err := s.JwtStorage.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, err
}
