package customers

import (
	"context"
	"time"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	Collection *mongo.Collection
}

// SignIn ...
func (c *Customer) SignIn(ctx context.Context, auth *models.SignInput) (string, error) {
	customer, _, err := c.findByUsername(ctx, auth.Username)
	if err != nil {
		return "", err
	}
	return createJWT(customer, auth)
}

// SignUp ...
func (c *Customer) SignUp(ctx context.Context, customer *models.Customer) error {
	_, isNot, err := c.findByUsername(ctx, customer.Username)
	if isNot {
		customer.RegistrationDate = time.Now()
		customer.ReleaseDate = time.Now()
		hash, err := encryptString(customer.Password)
		if err != nil {
			return err
		}
		customer.Password = hash
		return c.create(ctx, customer)
	}
	if err != nil {
		return err
	}
	return ErrUsernameAlready
}

// Profile ...
func (c *Customer) Profile(ctx context.Context, token string) error {
	return nil
}
