package stores

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/models"
)

// CustomerRepository ...
type CustomerRepository interface {
	Find(context.Context, *models.Customer) error
	FindAndCreate(context.Context, *models.Customer) error
	Update(context.Context, *models.Customer, interface{}) error
	FindAndUpdate(context.Context, *models.Customer) error
}
