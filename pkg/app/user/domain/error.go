package domain

import appError "Project/internal/error"

const (
	diferentPasswordErrorCode  = "password_diferents"
	emailAlreadyExistErrorCode = "email_exist"
	databaseErrorCode          = "database_error"
	serverErrorCode            = "server_error"
	emailOrPasswordIsWrongCode = "email_or_password_is_wrong"
)

const (
	diferentPasswordErrorMessage  = "the passwords are diferents"
	emailAlreadyExistErrorMessage = "the email already exist"
	databaseErrorMessage          = "the database had a error"
	serverErrorMessage            = "server had a error"
	emailOrPasswordIsWrongMessage = "email or password is wrong"
)

var DiferentPasswordError = appError.AppError{
	Code:    diferentPasswordErrorCode,
	Message: diferentPasswordErrorMessage,
}

var EmailAlreadyExistError = appError.AppError{
	Code:    emailAlreadyExistErrorCode,
	Message: emailAlreadyExistErrorMessage,
}

var DatabaseError = appError.AppError{
	Code:    databaseErrorCode,
	Message: databaseErrorMessage,
}

var ServerError = appError.AppError{
	Code:    serverErrorCode,
	Message: serverErrorMessage,
}

var EmailOrPasswordIsWrong = appError.AppError{
	Code:    emailOrPasswordIsWrongCode,
	Message: emailOrPasswordIsWrongMessage,
}
