package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongos/customer"
)

// Store ...
type Store struct {
	Customer customer.CustomerRepository
	JWT      *jwt.Config
	Cache    cache.Config
}
