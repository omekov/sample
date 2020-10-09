package customers

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUsernameAlready    = errors.New("such username already exists")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
