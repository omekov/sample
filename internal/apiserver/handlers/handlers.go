package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/omekov/sample/internal/apiserver/models"
	"go.mongodb.org/mongo-driver/bson"
)

// signIn godoc
// @Summary Sign auth
// @Description Sign auth client the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signin body models.Customer true "SignIn auth"
// @Success 200 {string} Token "qwerty"
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /signin [post]
func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var customer *models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		return
	}
	pass := customer.Credential.Password
	err := s.Store.Customer.FindCustomer(r.Context(), customer)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, err)
		return
	}
	err = customer.ComparePassword(pass)
	if err != nil {
		s.error(w, r, http.StatusUnauthorized, err)
		return
	}
	updateRealseDate := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "releasedate", Value: time.Now()},
			},
		},
	}
	if err := s.Store.Customer.UpdateCustomer(r.Context(), customer, updateRealseDate); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
	token, err := customer.GenerateJWT(customer, s.TokenSecret)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
	s.respond(w, r, http.StatusOK, token)
	return
}

// signUp godoc
// @Summary Sign Up new customer
// @Description Sign Up new customer the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signup body models.Customer true "SignUp customer"
// @Success 201 {object} models.Customer
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
	if err := s.Store.Customer.CreateCustomer(r.Context(), customer); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	s.respond(w, r, http.StatusCreated, nil)
	return
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
func (s *Server) whoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*models.Claims))
	}
}

// whoami godoc
// @Summary Refresh token
// @Description whoami input header Authorization Bearer <token>, return refresh in Claims
// @Tags sign
// @Accept  json
// @Produce  json
// @Success 200 {string} string "token"
// @Failure 400 {object} models.Error
// @Failure 401 {object} models.Error
// @Failure 403 {object} models.Error
// @Failure 404 {object} models.Error
// @Failure 500 {object} models.Error
// @Security ApiKeyAuth
// @Router /api/refresh [post]
func (s *Server) refreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}
		splitted := strings.Split(tokenHeader, " ")
		var c models.Customer
		token, err := c.RefreshToken(splitted[1], s.TokenSecret)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}
		s.respond(w, r, http.StatusOK, token)
	}
}

/*

// getOrder godoc
// @Summary Get details for a given orderId
// @Description Get details of order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order"
// @Success 200 {object} Order
// @Router /orders/{id} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["id"]
	for _, order := range orders {
		if order.ID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

// updateOrder godoc
// @Summary Update order identified by the given orderId
// @Description Update the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be updated"
// @Param order body Order true "Update order"
// @Success 200 {object} Order
// @Router /orders/{id} [put]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["id"]
	for i, order := range orders {
		if order.ID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

// deleteOrder godoc
// @Summary Delete order identified by the given orderId
// @Description Delete the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be deleted"
// @Success 204 "No Content"
// @Router /orders/{id} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["id"]
	for i, order := range orders {
		if order.ID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

*/
