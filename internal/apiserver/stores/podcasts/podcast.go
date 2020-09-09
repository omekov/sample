package podcasts

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Podcast struct {
	Collection *mongo.Collection
}

func (s *Podcast) CreatePodcast(ctx context.Context, podcast *models.Podcast) error {
	_, err := s.Collection.InsertOne(ctx, podcast)
	if err != nil {
		return err
	}
	return nil
}

func (s *Podcast) GetAllPodcasts(ctx context.Context) (*[]models.Podcast, error) {
	var podcasts []models.Podcast
	cur, err := s.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		podcast := models.Podcast{}
		err := cur.Decode(&podcast)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, podcast)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return &podcasts, nil
}
