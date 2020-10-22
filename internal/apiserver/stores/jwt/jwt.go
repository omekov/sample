package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
)

// Claims ...
type Claims struct {
	Customer models.Customer
	jwt.StandardClaims
}

// Config ...
type Config struct {
	TokenSecret []byte
	TokenMinute time.Duration
}

// NewConfig ...
func NewConfig(secret []byte, min time.Duration) *Config {
	return &Config{
		TokenSecret: secret,
		TokenMinute: min,
	}
}

var (
	errInvalidAccessToken = errors.New("invalid access token")
)

// NewJWT ...
func (c *Config) NewJWT(customer *models.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(c.TokenMinute * time.Minute).Unix(),
		},
	})
	return token.SignedString(c.TokenSecret)
}

// ParseJWT ...
func (c *Config) ParseJWT(tokenString string, claims *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidAccessToken
		}
		return c.TokenSecret, nil
	})
}

// Claims ...
func (c *Config) Claims(token string) (*Claims, error) {
	var claims Claims
	t, err := c.ParseJWT(token, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errInvalidAccessToken
	}
	return &claims, nil
}

// RefreshJWT ...
func (c *Config) RefreshJWT(token string) (string, error) {
	var claims Claims
	t, err := c.ParseJWT(token, &claims)
	if err != nil {
		return "", err
	}
	if !t.Valid {
		return "", errInvalidAccessToken
	}
	newToken, err := c.NewJWT(&claims.Customer)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
