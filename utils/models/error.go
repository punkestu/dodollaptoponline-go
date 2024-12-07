package models

import "strconv"

type Error struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func NewError(message string, code int) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func ToError(err error) *Error {
	newErr := err.(*Error)
	return newErr
}

func (e *Error) Error() string {
	return "Error " + strconv.Itoa(e.Code) + " : " + e.Message
}
