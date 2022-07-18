package routes

import (
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/handlers"
	"github.com/istefanini/restapi-beta/middleware"
)

func RegisteredRoutes(r *mux.Router) {
	r.Use(requestIDHandler)
	r.HandleFunc("/payment/api/v1/payments", handlers.GetPayments).Methods("GET")
}
