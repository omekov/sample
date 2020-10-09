package helpers

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	TokenSecret []byte
}

func (c *Config) GenerateJWT(customer *models.Customer, auth *models.SignInput) (string, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(auth.Password)); err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	customer.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})
	return token.SignedString(c.TokenSecret)
}
func (c *Config) EncryptString(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Config) ParseToken(tokenString string, claims *models.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidAccessToken
		}
		return c.TokenSecret, nil
	})
}

func (c *Config) ParseClaims(token *jwt.Token) *models.Claims {
	var result models.Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		result.Customer.Username = claims["Customer"].(string)
		result.Customer.FirstName = claims["firstname"].(string)
		result.Customer.RegistrationDate = claims["registrationDate"].(time.Time)
		result.ExpiresAt = claims["ExpiresAt"].(int64)
	}
	return &result
}
