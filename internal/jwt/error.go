package jwt

import appError "Project/internal/error"

const (
	tokenExpireErrorCode = "token_expired"
	tokenIsNotValidErrorCode = "token_is_not_valid"

)

const (
	tokenExpireErrorMessage = "the token expired"
	tokenIsNotValidErrorMessage = "token is not valid"
)

var TokenExpireError = appError.AppError{
	Code:    tokenExpireErrorCode,
	Message: tokenExpireErrorMessage,
}

var TokenIsNotValidError = appError.AppError{
	Code:    tokenIsNotValidErrorCode,
	Message: tokenIsNotValidErrorMessage,
}