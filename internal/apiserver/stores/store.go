package stores

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Store ...
type Store struct {
	AuthCollection    *mongo.Collection
	PodcastCollection *mongo.Collection
}

var jwtSecretKey = []byte("secret")

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

// SignIn ...
func (s *Store) SignIn(ctx context.Context, auth *models.SignInput) (string, error) {
	var customer *models.Customer
	err := s.AuthCollection.FindOne(ctx, bson.D{
		{Key: "username", Value: auth.Username},
	}).Decode(&customer)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(auth.Password))
	if err != nil {
		return "", err
	}
	return s.createJWT(customer)

}

func (s *Store) createJWT(customer *models.Customer) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})
	return token.SignedString([]byte("secret"))
}

func (s *Store) verifyToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Store) SignUp(ctx context.Context, customer *models.Customer) error {
	_, err := s.AuthCollection.InsertOne(ctx, bson.D{
		{Key: "username", Value: customer.Username},
		{Key: "password", Value: customer.Username},
		{Key: "firstname", Value: customer.FirstName},
		{Key: "registrationDate", Value: time.Now()},
		{Key: "releaseDate", Value: time.Now()},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreatePodcast(ctx context.Context, podcast *models.Podcast) error {
	result, err := s.PodcastCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: podcast.Title},
		{Key: "author", Value: podcast.Author},
	})
	fmt.Printf("Collection Result - %v ", result)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetAllPodcasts(ctx context.Context) (*[]models.Podcast, error) {
	var podcasts []models.Podcast
	cur, err := s.PodcastCollection.Find(ctx, bson.D{})
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
