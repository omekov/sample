package stores

import (
	"context"
	"time"

	"github.com/omekov/sample/internal/apiserver/stores/customers"
	"github.com/omekov/sample/internal/apiserver/stores/podcasts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store ...
type Store struct {
	Customers customers.Customer
	Podcasts  podcasts.Podcast
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
