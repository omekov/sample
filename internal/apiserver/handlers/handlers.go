package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/omekov/sample/internal/apiserver/models"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
)

// signIn godoc
// @Summary Sign auth
// @Description Sign auth client the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signin body models.Credential true "SignIn auth"
// @Success 200 {object} models.Token
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /signin [post]
func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var credential *models.Credential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}
	customer := &models.Customer{
		Username:    credential.Username,
		ReleaseDate: time.Now(),
	}
	err := s.Store.CustomerRepo.FindAndUpdate(r.Context(), customer)
	if err != nil {
		s.error(w, r, http.StatusForbidden, errIncorrectEmailPassword)
		return
	}
	err = customer.ComparePassword(credential.Password)
	if err != nil {
		s.error(w, r, http.StatusForbidden, errIncorrectEmailPassword)
		return
	}
	accToken, err := s.Store.JWT.NewAccessJWT(customer)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	refToken, err := s.Store.JWT.NewRefreshJWT(customer)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := s.Store.Redis.SetCustomerIDAndRefreshToken(customer, refToken); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusOK, models.Token{
		AccessToken:  accToken,
		Refreshtoken: refToken,
	})
}

// signUp godoc
// @Summary Sign Up new customer
// @Description Sign Up new customer the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signup body models.Customer true "SignUp customer"
// @Success 201 {string} string	"ok"
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /signup [post]
func (s *Server) signUp(w http.ResponseWriter, r *http.Request) {
	var customer *models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}
	if err := s.Store.CustomerRepo.FindAndCreate(r.Context(), customer); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := s.Store.RabbitMQ.Send(customer, uuid.New()); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusCreated, nil)
}

// whoami godoc
// @Summary Parse token
// @Description whoami input header Authorization Bearer <token>, return parse in Claims
// @Tags sign
// @Accept  json
// @Produce  json
// @Header 200 {string} Authorization "Bearer token"
// @Success 200 {string} string {"Customer":{"username": "example@gmail.com", "firstname": "Adam" },"exp": 1602666876}
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Security ApiKeyAuth
// @Router /api/whoami [get]
func (s *Server) whoAmi(w http.ResponseWriter, r *http.Request) {
	s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*jwt.Claims))
}

// refresh godoc
// @Summary Refresh token
// @Description http body refreshtoken sign new refresh token
// @Tags sign
// @Accept  json
// @Produce  json
// @Param refresh body models.Token true "Refresh auth"
// @Success 200 {object} models.Token
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /refresh [post]
func (s *Server) refreshToken(w http.ResponseWriter, r *http.Request) {
	var token *models.Token
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}
	newToken, err := s.Store.JWT.GetRefreshJWT(token.Refreshtoken)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusOK, newToken)
}

func (s *Server) shutdown(w http.ResponseWriter, r *http.Request) {
	status := models.ServerStatus{
		ShutdownStatus: "On",
	}
	s.respond(w, r, http.StatusOK, status)

	if !atomic.CompareAndSwapUint32(&s.ReqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.ShutdownReq <- true
	}()
}
