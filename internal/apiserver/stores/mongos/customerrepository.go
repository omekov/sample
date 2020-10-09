package mongos

import (
	"context"
	"time"

	"github.com/omekov/sample/internal/apiserver/models"
	"github.com/omekov/sample/internal/apiserver/stores/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Customer ...
type Customer struct {
	DB     Database
	Helper helpers.Config
}

const collectionName = "customers"

// FindCustomer ...
func (c *Customer) FindCustomer(ctx context.Context, auth *models.SignInput) (string, error) {
	var customer models.Customer
	err := c.DB.Collection(collectionName).FindOne(ctx, bson.D{
		{Key: "username", Value: auth.Username},
	}).Decode(&customer)
	if err != nil {
		return "", err
	}
	updateRealseDate := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "releasedate", Value: time.Now()},
			},
		},
	}
	_, err = c.DB.Collection(collectionName).UpdateOne(ctx, customer.ID, updateRealseDate)
	if err != nil {
		return "", err
	}
	return c.Helper.GenerateJWT(&customer, auth)
}

// CreateCustomer ...
func (c *Customer) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	err := c.DB.Collection(collectionName).FindOne(ctx, bson.D{
		{Key: "username", Value: customer.Username},
	}).Decode(&customer)
	if err == mongo.ErrNoDocuments {
		hash, err := c.Helper.EncryptString(customer.Password)
		if err != nil {
			return err
		}
		customer.ID = primitive.NewObjectID()
		customer.Password = hash
		customer.ReleaseDate = time.Now()
		customer.RegistrationDate = time.Now()
		_, err = c.DB.Collection(collectionName).InsertOne(ctx, customer)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	return helpers.ErrUsernameAlready
}

// Customer ...
func (c *Customer) Customer(ctx context.Context, splitted []string) (*models.Claims, error) {
	tokenPart := splitted[1]
	claims := models.Claims{}
	token, err := c.Helper.ParseToken(tokenPart, &claims)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, helpers.ErrInvalidAccessToken
	}
	return &claims, nil
}
