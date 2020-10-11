package customers

import (
	"context"
	"time"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Customer ...
type Customer struct {
	Collection  *mongo.Collection
	TokenSecret []byte
}

// SignIn ...
func (c *Customer) SignIn(ctx context.Context, auth *models.Auth) (string, error) {
	customer, _, err := c.findByUsername(ctx, auth.Username)
	if err != nil {
		return "", err
	}
	if err := c.updateRealseDate(ctx, customer.ID); err != nil {
		return "", err
	}
	return c.generateJWT(customer, auth)
}

// SignUp ...
func (c *Customer) SignUp(ctx context.Context, customer *models.Customer) error {
	_, isNot, err := c.findByUsername(ctx, customer.Username)
	if isNot {
		hash, err := encryptString(customer.Password)
		if err != nil {
			return err
		}
		customer.ID = primitive.NewObjectID()
		customer.Password = hash
		customer.ReleaseDate = time.Now()
		customer.RegistrationDate = time.Now()
		return c.create(ctx, customer)
	}
	if err != nil {
		return err
	}
	return ErrUsernameAlready
}

// WhoAmi ...
func (c *Customer) WhoAmi(ctx context.Context, splitted []string) (*models.Claims, error) {
	tokenPart := splitted[1]
	claims := models.Claims{}
	token, err := c.parseToken(tokenPart, &claims)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidAccessToken
	}
	return &claims, nil
}
