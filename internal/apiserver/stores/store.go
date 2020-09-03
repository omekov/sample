package stores

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store ...
type Store struct {
	collection *mongo.Collection
}

// ConfigureStore ...
func (s *Store) ConfigureStore(ctx context.Context, URI string, Database string, Collection string) error {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	s.collection = client.Database(Database).Collection(Collection)
	return nil
}

// Repository ...
type Repository interface {
	Sign() Signer
}
