package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Customer Customer
	jwt.StandardClaims
}
type SignInput struct {
	Password string `json:"password,omitempty" example:"123456"`
	Username string `json:"username,omitempty" example:"example@gmail.com"`
}

type Customer struct {
	Username         string    `json:"username,omitempty" example:"example@gmail.com"`
	FirstName        string    `json:"firstname,omitempty" example:"Adam"`
	Password         string    `json:"password,omitempty" example:"123456"`
	Blocked          bool      `json:"blocked,omitempty" example:"false"`
	RegistrationDate time.Time `json:"registrationDate,omitempty" example:"2020-09-09T21:21:46+00:00"`
	ReleaseDate      time.Time `json:"releaseDate,omitempty" example:"2020-09-09T22:21:46+00:00"`
}

type Podcast struct {
	Title  string `json:"title,omitempty" example:"title"`
	Author string `json:"author,omiempty" example:"example@gmail.com"`
}
