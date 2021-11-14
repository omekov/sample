package apiserver

import (
	"context"
	"fmt"
	"log"
	gohttp "net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/omekov/sample/internal/apiserver/delivery/http"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/usecase"
	"github.com/omekov/sample/internal/config"
	"github.com/omekov/sample/pkg/logger"
)

// Run ...
func Run() {
	if err := APIServer(); err != nil {
		log.Fatal(err)
	}
}

func APIServer() error {
	cfg := config.Get()
	_log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	store := stores.NewStore(ctx, cfg, _log)

	_useCase := usecase.NewUseCase(cfg, store)

	server := http.Server{
		Logger: _log,
		Router: mux.NewRouter(),
		Server: gohttp.Server{
			Addr:         fmt.Sprintf(":%v", cfg.Port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		ShutdownReq: make(chan bool),
		UseCase:     _useCase,
	}
	server.Run()
	return nil
}
