package stores

import (
	"context"
	"time"

	"github.com/omekov/sample/internal/apiserver/stores/mongos"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store ...
type Store struct {
	Customer mongos.Customer
}

// ConfigureStore ...
func ConfigureStore(URI string, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
