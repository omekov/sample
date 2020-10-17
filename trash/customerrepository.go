package mongos

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"errors"
)

type config struct {
	DB     Database
	Collection string
}

// CustomerRepository ...
type CustomerRepository interface {
	FindCustomer(context.Context, *models.Customer) error
	CreateCustomer(context.Context, *models.Customer) error
	UpdateCustomer(context.Context, *models.Customer, interface{}) error
}

// NewCustomerRepository ...
func NewCustomerRepository(db Database, collection string) CustomerRepository {
	return &config{
		DB:     db,
		Collection: collection,
	}
}

// FindCustomer ...
func (c *config) FindCustomer(ctx context.Context, customer *models.Customer) error {
	err := c.DB.Collection(c.Collection).FindOne(ctx, bson.D{
		{Key: "username", Value: customer.Username},
	}).Decode(&customer)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCustomer ...
func (c *config) UpdateCustomer(ctx context.Context,  customer *models.Customer, update interface{}) error {
	if _, err := c.DB.Collection(c.Collection).UpdateOne(ctx, customer.ID, update); err != nil {
		return err
	}
	return nil
}

// CreateCustomer ...
func (c *config) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	if err := customer.Validate(); err != nil {
		return err
	}
	if err := customer.BeforeCreate(); err != nil {
		return err
	}
	customer.Sanitize()
	err := c.FindCustomer(ctx, customer)
	if err == mongo.ErrNoDocuments {
		_, err = c.DB.Collection(c.Collection).InsertOne(ctx, customer)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return errors.New("such username already exists")
}

