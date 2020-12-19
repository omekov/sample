package apiserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/omekov/sample/configs"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"
	log "github.com/sirupsen/logrus"
)

// Run ...
func Run() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

func app() error {
	env := configs.NewENV()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s := stores.Store{}
	store, err := s.New(ctx, env)
	if err != nil {
		return err
	}
	server := handlers.Server{
		Logger: log.New(),
		Router: mux.NewRouter(),
		Server: http.Server{
			Addr:         fmt.Sprintf(":%v", env.Port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		ShutdownReq: make(chan bool),
		Store:       store,
	}
	server.Run()
	return nil
}
