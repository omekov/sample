package models

import (
	"time"
	validation "github.com/go-ozzo/ozzo-validation"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
)
var (
	errUserNotFound       = errors.New("user not found")
	errUsernameAlready    = errors.New("such username already exists")
	errInvalidAccessToken = errors.New("invalid access token")
)

// Claims ...
type Claims struct {
	Customer Customer
	jwt.StandardClaims
}

// Customer ...
type Customer struct {
	ID               primitive.ObjectID `json:"-" swaggerignore:"true"`
	Username         string             `json:"username,omitempty" example:"example@gmail.com"`
	FirstName        string             `json:"firstname,omitempty" example:"Adam"`
	Password         string             `json:"password,omitempty" example:"123456"`
	RepeatPassword   string             `json:"repeatPassword,omitempty" example:"123456"`
	EncryptedPassword   string          `json:"-" swaggerignore:"true"`
	Blocked          bool               `json:"-" swaggerignore:"true"`
	Actived          bool               `json:"-" swaggerignore:"true"`
	RegistrationDate time.Time          `json:"-" swaggerignore:"true"`
	ReleaseDate      time.Time          `json:"-" swaggerignore:"true"`
}


// Validate ...
func (c *Customer) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Username, validation.Required, is.Email),
		validation.Field(&c.Password, validation.By(requiredIf(c.EncryptedPassword == "")), validation.Length(6, 100)),
		validation.Field(&c.RepeatPassword, validation.Required, validation.By(repeatPassword(c.Password, c.RepeatPassword)), validation.Length(6, 100)),
	)
}

// BeforeCreate ...
func (c *Customer) BeforeCreate() error {
	if len(c.Password) > 0 {
		enc, err := encryptString(c.Password)
		if err != nil {
			return err
		}
		c.EncryptedPassword = enc
	}
	return nil
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

// GenerateJWT ...
func (c *Customer) GenerateJWT(customer *Customer, tokenSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	})
	return token.SignedString(tokenSecret)
}

// ParseToken ...
func (c *Customer) ParseToken(tokenString string, claims *Claims, tokenSecret []byte) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errInvalidAccessToken
		}
		return tokenSecret, nil
	})
}

// ParseClaims ...
// func (c *Customer) ParseClaims(token *jwt.Token) *Claims {
// 	var result Claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		result.Customer.Username = claims["Customer"].(string)
// 		result.Customer.FirstName = claims["firstname"].(string)
// 		result.Customer.RegistrationDate = claims["registrationDate"].(time.Time)
// 		result.ExpiresAt = claims["ExpiresAt"].(int64)
// 	}
// 	return &result
// }

// Customer ...
func (c *Customer) Customer(splitted []string, tokenSecret []byte) (*Claims, error) {
	tokenPart := splitted[1]
	claims := Claims{}
	token, err := c.ParseToken(tokenPart, &claims, tokenSecret)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errInvalidAccessToken
	}
	return &claims, nil
}

// RefreshToken ...
func (c *Customer) RefreshToken(token string, tokenSecret []byte) (string, error) {
	claims := Claims{}
	tkn, err := c.ParseToken(token, &claims, tokenSecret)
	if err != nil {
		return "", err
	}
	if !tkn.Valid {
		return "", errInvalidAccessToken
	}
	newToken, err := c.GenerateJWT(&claims.Customer, tokenSecret)
	if err != nil {
		return "", err
	}
	return newToken, nil
}