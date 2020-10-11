package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
)

// Store ...
type Store struct {
	Customer mongos.CustomerRepository
}
