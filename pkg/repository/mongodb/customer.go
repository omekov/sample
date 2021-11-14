package mongodb
/*
import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


import (
	"context"

	"errors"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CUSTOMERSCOLLECTION = "customers"
	ROLESCOLLECTION     = "roles"
)

// CustomerRepository ...
type CustomerRepository struct {
	db Database
}

// NewCustomerRepository ...
func NewCustomerRepository(db Database) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Find(ctx context.Context, customer *models.Customer) error {
	return r.db.Collection(CUSTOMERSCOLLECTION).FindOne(
		ctx,
		bson.M{
			"username": customer.Username,
		},
	).Decode(&customer)
}

// FindAndUpdate ...
func (r *CustomerRepository) FindAndUpdate(ctx context.Context, customer *models.Customer) error {
	return r.db.Collection(CUSTOMERSCOLLECTION).FindOneAndUpdate(
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
func (r *CustomerRepository) Update(ctx context.Context, customer *models.Customer, update interface{}) error {
	if _, err := r.db.Collection(CUSTOMERSCOLLECTION).UpdateOne(ctx, customer.ID, update); err != nil {
		return err
	}
	return nil
}

// Create ...
func (r *CustomerRepository) FindAndCreate(ctx context.Context, customer *models.Customer) error {
	if err := customer.Validate(); err != nil {
		return err
	}
	if err := customer.BeforeCreate(); err != nil {
		return err
	}
	customer.Sanitize()
	err := r.Find(ctx, customer)
	if err == mongo.ErrNoDocuments {
		_, err = r.db.Collection(CUSTOMERSCOLLECTION).InsertOne(ctx, customer)
		if err != nil {
			return err
		}
		_, err = r.db.Collection(ROLESCOLLECTION).InsertOne(ctx, customer.Roles)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return errors.New("such username already exists")
}
*/