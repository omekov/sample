package stores

import "sample/internal/apiserver/models"

// Signer ...
type Signer interface {
	In(auth models.Auth)
	Up(customer models.Customer)
	Customer(token string)
}
