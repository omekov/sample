package contant

import (
	"errors"
	"fmt"
)

var (
	ErrIncorrectEmailPassword = errors.New("incorrect Username or Password")
	ErrNotAuthenticated       = errors.New("not authenticated")
	ErrInvalidAccessToken     = errors.New("invalid access token")
	ErrUserNotFound           = errors.New("user not found")
	ErrUsernameAlready        = errors.New("such username already exists")
	ErrFilterKeyPageIsExist   = errors.New("filter['page'] must not be empty")
)

func ErrStconvAtoi(field string, err error) error {
	return fmt.Errorf("field: %s %v", field, err)
}
