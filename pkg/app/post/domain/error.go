package domain

import appError "Project/internal/error"

const (
	youHaveNotAccessErrorCode  = "user_do_not_access"
	emailAlreadyExistErrorCode = "email_exist"
	databaseErrorCode          = "database_error"
	serverErrorCode            = "server_error"
)

const (
	youHaveNotAccessErrorMessage  = "you do not access so this post "
	emailAlreadyExistErrorMessage = "the email already exist"
	databaseErrorMessage          = "the database had a error"
	serverErrorMessage            = "server had a error"
)

var YouHaveNotAccessError = appError.AppError{
	Code:    youHaveNotAccessErrorCode,
	Message: youHaveNotAccessErrorMessage,
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
