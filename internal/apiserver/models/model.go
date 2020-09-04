package models

import (
	"time"
)

// Order represents the model for an order
type Order struct {
	ID           string     `json:"id,omitempty" example:"1"`
	CustomerName string     `json:"customerName,omitempty" example:"Leo Messi"`
	OrderedAt    *time.Time `json:"orderedAt,omitempty" example:"2020-09-09T21:21:46+00:00"`
	Items        []Item     `json:"items,omitempty"`
}

// Item represents the model for an item in the order
type Item struct {
	ID          string `json:"id,omitempty" example:"A1B2C3"`
	Description string `json:"description,omitempty" example:"A random description"`
	Quantity    int    `json:"quantity,omitempty" example:"1"`
}

type SignInput struct {
	Password string `json:"password,omitempty" example:"123456"`
	Username string `json:"username,omitempty" example:"example@gmail.com"`
}

type Customer struct {
	Username         string    `json:"username,omitempty" example:"example@gmail.com"`
	FirstName        string    `json:"firstname,omitempty" example:"Adam"`
	Password         string    `json:"password,omitempty" example:"123456"`
	RegistrationDate time.Time `json:"registrationDate,omitempty" example:"2020-09-09T21:21:46+00:00"`
	ReleaseDate      time.Time `json:"releaseDate,omitempty" example:"2020-09-09T22:21:46+00:00"`
}
