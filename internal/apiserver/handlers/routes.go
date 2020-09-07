package handlers

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/omekov/sample/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/omekov/sample/internal/apiserver/stores"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Server ...
type Server struct {
	Router *mux.Router
	Logger *logrus.Logger
	Store  *stores.Store
}

// ConfigureRouter ...
// @title Sample API
// @version 1.0
// @description This is a sample service for managment
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email umekovazamat@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /
func (s *Server) ConfigureRouter(PORT string) *mux.Router {
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.Router.Use(s.logRequest)
	s.Router.HandleFunc("/signin", s.signIn).Methods(http.MethodPost)
	s.Router.HandleFunc("/signup", s.signUp).Methods(http.MethodPost)
	s.Router.HandleFunc("/profile", s.profile).Methods(http.MethodGet)
	s.Router.HandleFunc("/podcasts", s.createPodcast).Methods(http.MethodPost)
	s.Router.HandleFunc("/podcasts", s.getPodcasts).Methods(http.MethodGet)

	// Swagger
	s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), s.Router))
	return s.Router
}
