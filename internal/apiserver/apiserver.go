package apiserver

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/models"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/stores/helpers"
	"github.com/omekov/sample/internal/apiserver/stores/mongos"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	MONGOURI               = "MONGOURI"
	MONGONAME              = "MONGONAME"
	MONGOUSERNAME          = "MONGOUSERNAME"
	MONGOPASSWORD          = "MONGOPASSWORD"
	MONGOAUTHCOLLECTION    = "MONGOAUTHCOLLECTION"
	MONGOPODCASTCOLLECTION = "MONGOPODCASTCOLLECTION"
	TOKENSECRET            = "TOKENSECRET"
	PORT                   = "PORT"
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

// func connections() {
// 	FlagAndLoadENV()
// 	db, err := stores.ConfigureStore(IsReadyENV(MONGOURI), IsReadyENV(MONGONAME))
// 	if err != nil {
// 		log.Fatal("mongodb error ", err)
// 	}
// 	fmt.Printf("Connection db - %s", db.Name())
// 	server := handlers.Server{
// 		Router: mux.NewRouter(),
// 		Logger: logrus.New(),
// 		Store: &stores.Store{
// 			Customers: customers.Customer{
// 				Collection:  db.Collection(IsReadyENV(MONGOAUTHCOLLECTION)),
// 				TokenSecret: []byte(IsReadyENV(TOKENSECRET)),
// 			},
// 		},
// 	}
// 	server.ConfigureRouter(IsReadyENV(PORT))
// }

// Run ...
func Run() {
	if err := newServer(); err != nil {
		log.Fatal("mongodb error ", err)
	}
}

// getConfig ...
func getConfig() *models.MongoConfig {
	return &models.MongoConfig{
		Username:     IsReadyENV(MONGOUSERNAME),
		Password:     IsReadyENV(MONGOPASSWORD),
		Url:          IsReadyENV(MONGOURI),
		DatabaseName: IsReadyENV(MONGONAME),
	}
}

func newServer() error {
	FlagAndLoadENV()
	dbClient, err := mongos.NewClient(getConfig())
	if err != nil {
		return err
	}
	if err = dbClient.Connect(); err != nil {
		return err
	}
	db := mongos.NewDatabase(getConfig(), dbClient)
	server := handlers.Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store: &stores.Store{
			Customer: mongos.Customer{
				DB: db,
				Helper: helpers.Config{
					TokenSecret: []byte(IsReadyENV(TOKENSECRET)),
				},
			},
		},
	}
	server.ConfigureRouter(IsReadyENV(PORT))
	return nil
}
