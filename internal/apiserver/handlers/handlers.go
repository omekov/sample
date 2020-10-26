package handlers

import (
	"encoding/json"
	"net/http"
	"time"

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
func (s *Server) signIn() http.HandlerFunc {
	var credential *models.Credential
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		customer := &models.Customer{
			Username:    credential.Username,
			ReleaseDate: time.Now(),
		}
		err := s.Store.Customer.FindAndUpdate(r.Context(), customer)
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
		if err := s.Store.Cache.SetCustomerIDAndRefreshToken(customer, refToken); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, models.Token{
			AccessToken:  accToken,
			Refreshtoken: refToken,
		})
		return
	}
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
func (s *Server) signUp() http.HandlerFunc {
	var customer *models.Customer
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := s.Store.Customer.Create(r.Context(), customer); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusCreated, nil)
		return
	}
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
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*jwt.Claims))
	}
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
func (s *Server) refreshToken() http.HandlerFunc {
	var token *models.Token
	return func(w http.ResponseWriter, r *http.Request) {
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
