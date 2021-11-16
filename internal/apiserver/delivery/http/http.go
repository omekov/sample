package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/omekov/sample/internal/apiserver/usecase"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server ...
type Server struct {
	Logger  *logrus.Logger
	UseCase *usecase.UseCase
	Router  *mux.Router
	http.Server
	ShutdownReq chan bool
	ReqCount    uint32
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Handlers ...
func (s *Server) handlers() *mux.Router {
	s.Router.Use(s.setRequestID)
	s.Router.Use(s.setHeaderAccessControlAllow)
	s.Router.Use(s.logRequest)
	s.Router.HandleFunc("/signin", s.signIn).Methods(http.MethodPost, http.MethodOptions)
	// s.Router.HandleFunc("/signup", s.signUp).Methods(http.MethodPost, http.MethodOptions)
	// s.Router.HandleFunc("/refresh", s.refreshToken).Methods(http.MethodPost, http.MethodOptions)
	// s.Router.HandleFunc("/shutdown", s.shutdown).Methods(http.MethodGet, http.MethodOptions)
	private := s.Router.PathPrefix("/api").Subrouter()
	private.Use(s.authenticateUser)
	// private.HandleFunc("/whoami", s.whoAmi).Methods(http.MethodGet, http.MethodOptions)

	// Swagger
	s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return s.Router
}

func (s *Server) waitShutdown() {
	signalChan := make(chan os.Signal)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	select {
	case sig := <-signalChan:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.ShutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stoping http server ...")
	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	}
	os.Exit(0)
}

// Run ...
func (s *Server) Run() {
	s.Handler = s.handlers()
	done := make(chan bool)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
		done <- true
	}()

	//wait shutdown
	s.waitShutdown()

	<-done
}
