package models

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	errUserNotFound    = errors.New("user not found")
	errUsernameAlready = errors.New("such username already exists")
)

// Credential ...
type Credential struct {
	Username string `bson:"username" json:"username,omitempty" example:"example@gmail.com"`
	Password string `bson:"password" json:"password,omitempty" example:"-"`
}

// Customer ...
type Customer struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-" swaggerignore:"true"`
	Username          string             `bson:"username" json:"username,omitempty" example:"example@gmail.com"`
	Password          string             `bson:"password" json:"password,omitempty" example:"-"`
	FirstName         string             `bson:"firstname,omitempty" json:"firstname,omitempty" example:"Adam"`
	RepeatPassword    string             `json:"repeatPassword,omitempty" example:"123456"`
	EncryptedPassword string             `bson:"encryptedPassword,omitempty" json:"-" swaggerignore:"true"`
	Blocked           bool               `bson:"blocked,omitempty" json:"-" swaggerignore:"true"`
	Actived           bool               `bson:"actived,omitempty" json:"-" swaggerignore:"true"`
	RegistrationDate  time.Time          `bson:"registrationDate,omitempty" json:"-" swaggerignore:"true"`
	ReleaseDate       time.Time          `bson:"releaseDate,omitempty" json:"-" swaggerignore:"true"`
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
