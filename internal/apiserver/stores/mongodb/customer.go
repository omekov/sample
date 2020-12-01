package mongodb

import (
	"context"

	"errors"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type config struct {
	DB         Database
	Collection string
}

// CustomerRepository ...
type CustomerRepository interface {
	Find(context.Context, *models.Customer) error
	Create(context.Context, *models.Customer) error
	Update(context.Context, *models.Customer, interface{}) error
	FindAndUpdate(context.Context, *models.Customer) error
}

// NewCustomerRepository ...
func NewCustomerRepository(db Database, collection string) CustomerRepository {
	return &config{
		DB:         db,
		Collection: collection,
	}
}

func (c *config) Find(ctx context.Context, customer *models.Customer) error {
	return c.DB.Collection(c.Collection).FindOne(
		ctx,
		bson.M{
			"username": customer.Username,
		},
	).Decode(&customer)
}

// FindAndUpdate ...
func (c *config) FindAndUpdate(ctx context.Context, customer *models.Customer) error {
	return c.DB.Collection(c.Collection).FindOneAndUpdate(
		ctx,
		bson.M{
			"username": customer.Username,
		},
		bson.D{
			{
				"$set",
				bson.M{
					"releaseDate": customer.ReleaseDate,
				},
			},
		},
	).Decode(&customer)
}

// Update ...
func (c *config) Update(ctx context.Context, customer *models.Customer, update interface{}) error {
	if _, err := c.DB.Collection(c.Collection).UpdateOne(ctx, customer.ID, update); err != nil {
		return err
	}
	return nil
}

// Create ...
func (c *config) Create(ctx context.Context, customer *models.Customer) error {
	if err := customer.Validate(); err != nil {
		return err
	}
	if err := customer.BeforeCreate(); err != nil {
		return err
	}
	customer.Sanitize()
	err := c.Find(ctx, customer)
	if err == mongo.ErrNoDocuments {
		_, err = c.DB.Collection(c.Collection).InsertOne(ctx, customer)
		if err != nil {
			return err
		}
		_, err = c.DB.Collection("roles").InsertOne(ctx, customer.Roles)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return errors.New("such username already exists")
}
