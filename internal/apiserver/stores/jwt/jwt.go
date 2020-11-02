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
	RefreshTokenSecret []byte
	AccessTokenSecret  []byte
}

var (
	errInvalidAccessToken = errors.New("invalid access token")
)

// NewRefreshJWT ...
func (c *Config) NewRefreshJWT(customer *models.Customer) (string, error) {
	return c.newJWT(customer, 10, c.RefreshTokenSecret)
}

// NewAccessJWT ...
func (c *Config) NewAccessJWT(customer *models.Customer) (string, error) {
	return c.newJWT(customer, 7, c.AccessTokenSecret)
}

func (c *Config) newJWT(customer *models.Customer, tokentime time.Duration, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokentime * time.Minute).Unix(),
		},
	})
	return token.SignedString(secret)
}

// GetClaims ...
func (c *Config) GetClaims(token string) (*Claims, error) {
	var claims Claims
	t, err := c.parseJWT(token, c.AccessTokenSecret, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errInvalidAccessToken
	}
	return &claims, nil
}

// GetRefreshJWT ...
func (c *Config) GetRefreshJWT(refToken string) (*models.Token, error) {
	var claims Claims
	t, err := c.parseJWT(refToken, c.RefreshTokenSecret, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errInvalidAccessToken
	}
	newRefToken, err := c.NewRefreshJWT(&claims.Customer)
	if err != nil {
		return nil, err
	}
	newAccToken, err := c.NewAccessJWT(&claims.Customer)
	if err != nil {
		return nil, err
	}
	return &models.Token{
		AccessToken:  newAccToken,
		Refreshtoken: newRefToken,
	}, nil
}

func (c *Config) parseJWT(tokenString string, key []byte, claims *Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidAccessToken
		}
		return key, nil
	})
}

/*
func (c *Config) parseJWT(tokenString string, key []byte) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if token.Valid {
		if claims, ok := token.Claims.(Claims); ok {
			return &claims, nil
		}
		return nil, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
*/
