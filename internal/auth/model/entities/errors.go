package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user with this email or username already exists")
	ErrPasswordNotSecure = errors.New("provided password is not secure enough")
)
