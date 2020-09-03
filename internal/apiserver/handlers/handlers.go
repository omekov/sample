package handlers

import (
	"encoding/json"
	"net/http"
	"sample/internal/apiserver/models"
)

// signIn godoc
// @Summary Sign auth
// @Description Sign auth client the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signin body models.Auth true "SignIn auth"
// @Success 200 {object} models.Auth
// @Router /signin [post]
func (s *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var auth models.Auth
	json.NewDecoder(r.Body).Decode(&auth)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}

// signUp godoc
// @Summary Sign Up new customer
// @Description Sign Up new customer the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signup body models.Customer true "SignUp customer"
// @Success 201 {object} models.Customer
// @Router /signup [post]
func (s *Server) signUp(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
	// r.Response.StatusCode = http.StatusCreated
}

// profile godoc
// @Summary Profile customer
// @Description Profile customer the input paylod
// @Tags sign
// @Accept  json
// @Produce  json
// @Param signup body models.Customer true "Profile customer"
// @Success 200 {object} models.Customer
// @Router /profile [post]
func (s *Server) profile(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

/*
// createOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 201 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.ID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}


// getOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

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
