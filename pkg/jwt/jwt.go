package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/omekov/sample/internal/apiserver/delivery/http"
	"github.com/omekov/sample/internal/config"
	"github.com/omekov/sample/pkg/contant"
	"github.com/omekov/sample/pkg/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Claims ...
type Claims struct {
	Customer domain.User
	jwtgo.StandardClaims
}

// Config ...
type Config struct {
	refreshTokenSecret []byte
	accessTokenSecret  []byte
}

func NewJWT(cfg config.JWT) *Config {
	return &Config{
		refreshTokenSecret: []byte(cfg.Refresh),
		accessTokenSecret:  []byte(cfg.Access),
	}
}

// NewRefreshJWT ...
func (c *Config) NewRefreshJWT(customer *domain.User) (string, error) {
	return c.newJWT(customer, 10, c.refreshTokenSecret)
}

// NewAccessJWT ...
func (c *Config) NewAccessJWT(customer *domain.User) (string, error) {
	return c.newJWT(customer, 7, c.accessTokenSecret)
}

func (c *Config) newJWT(user *domain.User, tokentime time.Duration, secret []byte) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, Claims{
		Customer: *user,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(tokentime * time.Minute).Unix(),
		},
	})
	return token.SignedString(secret)
}

// GetClaims ...
func (c *Config) GetClaims(token string) (*Claims, error) {
	var claims Claims
	t, err := c.parseJWT(token, c.accessTokenSecret, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, contant.ErrInvalidAccessToken
	}
	return &claims, nil
}

// GetRefreshJWT ...
func (c *Config) GetRefreshJWT(refToken string) (*http.Token, error) {
	var claims Claims
	t, err := c.parseJWT(refToken, c.refreshTokenSecret, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, contant.ErrInvalidAccessToken
	}
	newRefToken, err := c.NewRefreshJWT(&claims.Customer)
	if err != nil {
		return nil, err
	}
	newAccToken, err := c.NewAccessJWT(&claims.Customer)
	if err != nil {
		return nil, err
	}
	return &http.Token{
		AccessToken:  newAccToken,
		Refreshtoken: newRefToken,
	}, nil
}

func (c *Config) parseJWT(tokenString string, key []byte, claims *Claims) (*http.Token, error) {
	return jwtgo.ParseWithClaims(tokenString, claims, func(token *http.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, contant.ErrInvalidAccessToken
		}
		return key, nil
	})
}

// Validate ...
func (c *Config) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Username, validation.Required, is.Email),
		validation.Field(&c.Password, validation.By(requiredIf(c.EncryptedPassword == "")), validation.Length(6, 100)),
		validation.Field(&c.RepeatPassword, validation.Required, validation.By(repeatPassword(c.Password, c.RepeatPassword)), validation.Length(6, 100)),
	)
}

// BeforeCreate ...
func (c *Config) BeforeCreate(password string) (enc string, err error) {
	if len(password) > 0 {
		enc, err = encryptString(password)
		if err != nil {
			return enc, err
		}
	}
	return enc, nil
}

func encryptString(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ComparePassword ...
func (c *Customer) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.EncryptedPassword), []byte(password))
}

// Sanitize ...
func (c *Customer) Sanitize() {
	c.ID = primitive.NewObjectID()
	c.Password = ""
	c.RepeatPassword = ""
	c.ReleaseDate = time.Now()
	c.RegistrationDate = time.Now()
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
