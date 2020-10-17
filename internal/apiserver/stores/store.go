package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/mongos/customer"
)

// Store ...
type Store struct {
	Customer customer.CustomerRepository
}
