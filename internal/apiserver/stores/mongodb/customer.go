package mongodb

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

type Repository struct {
	DB       Database
	Customer CustomerRepository
}

// CustomerRepository ...
type CustomerRepository interface {
	Find(context.Context, *models.Customer) error
	FindAndCreate(context.Context, *models.Customer) error
	Update(context.Context, *models.Customer, interface{}) error
	FindAndUpdate(context.Context, *models.Customer) error
}

func (r *Repository) Find(ctx context.Context, customer *models.Customer) error {
	return r.DB.Collection(CUSTOMERSCOLLECTION).FindOne(
		ctx,
		bson.M{
			"username": customer.Username,
		},
	).Decode(&customer)
}

// FindAndUpdate ...
func (r *Repository) FindAndUpdate(ctx context.Context, customer *models.Customer) error {
	return r.DB.Collection(CUSTOMERSCOLLECTION).FindOneAndUpdate(
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
func (r *Repository) Update(ctx context.Context, customer *models.Customer, update interface{}) error {
	if _, err := r.DB.Collection(CUSTOMERSCOLLECTION).UpdateOne(ctx, customer.ID, update); err != nil {
		return err
	}
	return nil
}

// Create ...
func (r *Repository) FindAndCreate(ctx context.Context, customer *models.Customer) error {
	if err := customer.Validate(); err != nil {
		return err
	}
	if err := customer.BeforeCreate(); err != nil {
		return err
	}
	customer.Sanitize()
	err := r.Find(ctx, customer)
	if err == mongo.ErrNoDocuments {
		_, err = r.DB.Collection(CUSTOMERSCOLLECTION).InsertOne(ctx, customer)
		if err != nil {
			return err
		}
		_, err = r.DB.Collection(ROLESCOLLECTION).InsertOne(ctx, customer.Roles)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return errors.New("such username already exists")
}
