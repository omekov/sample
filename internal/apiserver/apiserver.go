package apiserver

import (
	"context"
	"log"
	"os"

	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	// MONGOURI ...
	MONGOURI string = ""
	// PORT ...
	PORT string = "80"
	// MONGOCOLLECTION ...
	MONGOCOLLECTION string = ""
	// MONGONAME ...
	MONGONAME string = ""
)

func connections() {
	ctx := context.TODO()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	MONGOURI = os.Getenv("MONGOURI")
	MONGONAME = os.Getenv("MONGONAME")
	MONGOCOLLECTION = os.Getenv("MONGOCOLLECTION")
	PORT = os.Getenv("PORT")
	server := handlers.Server{
		Router: mux.NewRouter(),
		Logger: logrus.New(),
		Store:  &stores.Store{},
	}
	if err := server.Store.ConfigureStore(ctx, MONGOURI, MONGONAME, MONGOCOLLECTION); err != nil {
		log.Fatal("mongodb error ", err)
	}
	server.ConfigureRouter(PORT)
}

// Run ...
func Run() {
	connections()
}
