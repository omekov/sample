package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoConfig ...
type MongoConfig struct {
	Username     string
	Password     string
	DatabaseName string
	Url          string
}

// Claims ...
type Claims struct {
	Customer Customer
	jwt.StandardClaims
}

// Auth ...
type Auth struct {
	Username string `json:"username,omitempty" example:"example@gmail.com"`
	Password string `json:"password,omitempty" example:"123456"`
}

// Customer ...
type Customer struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username         string             `json:"username,omitempty" bson:"username,omitempty" example:"example@gmail.com"`
	FirstName        string             `json:"firstname,omitempty" example:"Adam"`
	Password         string             `json:"password,omitempty" example:"123456"`
	Blocked          bool               `json:"blocked,omitempty" example:"false"`
	Actived          bool               `json:"actived,omitempty" example:"false"`
	RegistrationDate time.Time          `json:"registrationDate,omitempty" example:"2020-09-09T21:21:46+00:00"`
	ReleaseDate      time.Time          `json:"releaseDate,omitempty" example:"2020-09-09T22:21:46+00:00"`
}

// Error ...
type Error struct {
	Error string `json:"error,omitempty" example:"error"`
}
