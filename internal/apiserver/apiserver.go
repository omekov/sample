package apiserver

import (
	"log"

	"github.com/omekov/sample/config"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
	"github.com/omekov/sample/internal/apiserver/stores/mongos/customer"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Run ...
func Run() {
	if err := newServer(); err != nil {
		log.Fatal(err)
	}
}

func newServer() error {
	config.Init(".env")
	dbClient, err := mongos.NewClient(config.GetMongoConfig())
	if err != nil {
		return err
	}
	if err = dbClient.Connect(); err != nil {
		return err
	}
	redisClient, err := cache.NewClient(config.GetRedisConfig())
	if err != nil {
		return err
	}
	db := mongos.NewDatabase(config.GetMongoConfig(), dbClient)
	customer := customer.NewCustomerRepository(
		db,
		config.IsReadyENV(config.MONGOCUSTOMERSCOLLECTION),
	)
	server := handlers.Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store: &stores.Store{
			Customer: customer,
			JWT: jwt.Config{
				RefreshTokenSecret: []byte(config.IsReadyENV(config.REFRESHTOKENSECRET)),
				AccessTokenSecret:  []byte(config.IsReadyENV(config.ACCESSTOKENSECRET)),
			},
			Cache: cache.Config{
				Client: redisClient,
			},
		},
	}
	server.ConfigureRouter(config.IsReadyENV(config.PORT))
	return nil
}
