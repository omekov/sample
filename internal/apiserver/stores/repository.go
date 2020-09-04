package stores

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/models"
)

// CustomerRepositorer ...
type CustomerRepositorer interface {
	In(auth models.SignInput)
	Up(customer models.Customer)
	Customer(token string)
}

// UseCase ...
type UseCase interface {
	SignUp(ctx context.Context, username, passwword string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.Customer, error)
}
