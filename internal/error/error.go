package error

import "fmt"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
