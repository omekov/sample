package apiserver

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	// MONGOURI ...
	MONGOURI string = ""
	// PORT ...
	PORT string = "80"
	// MONGOAUTHCOLLECTION ...
	MONGOAUTHCOLLECTION string = ""
	// MONGOPODCASTCOLLECTION ...
	MONGOPODCASTCOLLECTION string = ""
	// MONGONAME ...
	MONGONAME string = ""
)

func connections() {
	useENV := flag.Bool("env", false, "not load env file")
	flag.Parse()
	if *useENV {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file ", err)
		}
	}

	MONGOURI = os.Getenv("MONGOURI")
	MONGONAME = os.Getenv("MONGONAME")
	MONGOAUTHCOLLECTION = os.Getenv("MONGOAUTHCOLLECTION")
	MONGOPODCASTCOLLECTION = os.Getenv("MONGOPODCASTCOLLECTION")
	PORT = os.Getenv("PORT")
	db, err := stores.ConfigureStore(MONGOURI, MONGONAME)
	if err != nil {
		log.Fatal("mongodb error ", err)
	}
	fmt.Printf("Connection db - %s", db.Name())
	server := handlers.Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store: &stores.Store{
			AuthCollection:    db.Collection(MONGOAUTHCOLLECTION),
			PodcastCollection: db.Collection(MONGOPODCASTCOLLECTION),
		},
	}
	server.ConfigureRouter(PORT)
}

// Run ...
func Run() {
	connections()
}
