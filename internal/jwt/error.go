package jwt

import appError "Project/internal/error"

const (
	tokenExpireErrorCode = "password_diferents"
)

const (
	tokenExpireErrorMessage = "the passwords are diferents"
)

var TokenExpireError = appError.AppError{
	Code:    tokenExpireErrorCode,
	Message: tokenExpireErrorMessage,
}
