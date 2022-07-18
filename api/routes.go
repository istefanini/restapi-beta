package api

import (
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/handlers"
	"github.com/istefanini/restapi-beta/main"
)

func RegisterRoutes(r *mux.Router) {
	r.Use(requestIDHandler)
	r.HandleFunc("/payment/api/v1/payments", handlers.GetPayments).Methods("GET")
}
