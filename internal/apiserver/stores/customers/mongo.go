package customers

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindByUsername ...
func (c *Customer) findByUsername(ctx context.Context, username string) (*models.Customer, bool, error) {
	var customer *models.Customer
	err := c.Collection.FindOne(ctx, bson.D{
		{Key: "username", Value: username},
	}).Decode(&customer)
	if err != nil {
		return nil, (err == mongo.ErrNoDocuments), err
	}
	return customer, false, nil
}

// Create ...
func (c *Customer) create(ctx context.Context, customer *models.Customer) error {
	_, err := c.Collection.InsertOne(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}
