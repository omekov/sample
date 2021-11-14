package contant

import "errors"

var (
	ErrIncorrectEmailPassword = errors.New("incorrect Username or Password")
	ErrNotAuthenticated       = errors.New("not authenticated")
	ErrInvalidAccessToken     = errors.New("invalid access token")
	ErrUserNotFound           = errors.New("user not found")
	ErrUsernameAlready        = errors.New("such username already exists")
)
