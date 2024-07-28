package error

import "errors"

var (
	DiferentPasswordError = errors.New("the passwords are diferents")
)