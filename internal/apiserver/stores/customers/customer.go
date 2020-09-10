package customers

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customer struct {
	Collection  *mongo.Collection
	TokenSecret []byte
}

// SignIn ...
func (c *Customer) SignIn(ctx context.Context, auth *models.SignInput) (string, error) {
	customer, _, err := c.findByUsername(ctx, auth.Username)
	if err != nil {
		return "", err
	}
	return c.generateJWT(customer, auth)
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
func (c *Customer) Profile(ctx context.Context, token string) (*models.Claims, error) {
	jwtToken, err := c.parseToken(token)
	if err != nil {
		return nil, err
	}
	return parseClaims(jwtToken), nil
}

// Profile ...
func (c *Customer) WhoAmi(ctx context.Context, splitted []string) (*models.Claims, error) {
	tokenPart := splitted[1] //Получаем вторую часть токена
	tk := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return c.TokenSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidAccessToken
	}

	//Всё прошло хорошо, продолжаем выполнение запроса
	fmt.Sprintf("Username -  %s", tk.Customer.Username) //Полезно для мониторинга
	return tk, nil
}
