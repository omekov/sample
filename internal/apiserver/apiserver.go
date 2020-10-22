package apiserver

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/models"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
	"github.com/omekov/sample/internal/apiserver/stores/mongos/customer"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	MONGOURI               = "MONGOURI" // MONGOURI - ...
	MONGONAME              = "MONGONAME"
	MONGOUSERNAME          = "MONGOUSERNAME"
	MONGOPASSWORD          = "MONGOPASSWORD"
	MONGOAUTHCOLLECTION    = "MONGOAUTHCOLLECTION"
	MONGOPODCASTCOLLECTION = "MONGOPODCASTCOLLECTION"
	TOKENSECRET            = "TOKENSECRET"
	PORT                   = "PORT"
	REDISPASSWORD          = "REDISPASSWORD"
	REDISURI               = "REDISURI"
)

// IsReadyENV ...
func IsReadyENV(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("Error is not env - %s", key)
		return ""
	}
	return os.Getenv(key)
}

// FlagAndLoadENV ...
func FlagAndLoadENV() {
	useENV := flag.Bool("env", false, "not load env file")
	flag.Parse()
	if *useENV {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file ", err)
		}
	}
}

// Run ...
func Run() {
	if err := newServer(); err != nil {
		log.Fatal(err)
	}
}

// GetMongoConfig ...
func GetMongoConfig() *models.MongoConfig {
	return &models.MongoConfig{
		Username:     IsReadyENV(MONGOUSERNAME),
		Password:     IsReadyENV(MONGOPASSWORD),
		URL:          IsReadyENV(MONGOURI),
		DatabaseName: IsReadyENV(MONGONAME),
	}
}

// GetRedisConfig ...
func GetRedisConfig() *models.RedisConfig {
	return &models.RedisConfig{
		Password: IsReadyENV(REDISPASSWORD),
		URL:      IsReadyENV(REDISURI),
	}
}

func newServer() error {
	FlagAndLoadENV()
	dbClient, err := mongos.NewClient(GetMongoConfig())
	if err != nil {
		return err
	}
	if err = dbClient.Connect(); err != nil {
		return err
	}
	redisClient, err := cache.NewClient(GetRedisConfig())
	if err != nil {
		return err
	}
	db := mongos.NewDatabase(GetMongoConfig(), dbClient)
	customer := customer.NewCustomerRepository(
		db,
		"customers",
	)
	jwtoken := jwt.NewConfig([]byte(IsReadyENV(TOKENSECRET)), 5)
	server := handlers.Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store: &stores.Store{
			Customer: customer,
			JWT:      jwtoken,
			Cache: cache.Config{
				Client: redisClient,
			},
		},
	}
	server.ConfigureRouter(IsReadyENV(PORT))
	return nil
}
