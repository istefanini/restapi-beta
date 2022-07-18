package api

import (
	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.Use(requestIDHandler)
	r.HandleFunc("/payment/api/v1/payments", GetPayments).Methods("GET")
	r.HandleFunc("/payment/api/v1/notificaction-mol-payment", CreatePayment).Methods("POST")
	r.HandleFunc("/payment/api/v1/payments/{external_reference}", GetOnePayment).Methods("GET")
}
