package stores

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Store ...
type Store struct {
	collection *mongo.Collection
}

// ConfigureStore ...
func (s *Store) ConfigureStore(ctx context.Context, URI string, Database string, Collection string) error {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
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

// SignIn ...
func (s *Store) SignIn(auth *models.SignInput) (string, error) {
	result := new(models.Customer)
	err := s.collection.FindOne(context.TODO(), bson.D{{"username", auth.Username}}).Decode(&result)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(result.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"firtsname": result.FirstName,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
