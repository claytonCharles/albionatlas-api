package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("User nor found")
	ErrInvalidCredentials = errors.New("Mail or Password invalid")
	ErrFail               = errors.New("Internal Error")
)
