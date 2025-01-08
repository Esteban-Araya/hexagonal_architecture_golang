package jwt

import (
	jwt_lib "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	User_id int `json:"user_id"`
	jwt_lib.RegisteredClaims
}

type RefreshAndAccessTokens struct {
	Refresh string `json:"refresh_token"`
	Access string `json:"access_token"`
}