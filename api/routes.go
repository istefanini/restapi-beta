package api

import (
	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	protected := r.NewRoute().Subrouter()
	protected.Use(requestIDHandler)
	protected.HandleFunc("/payment/api/v1/notificaction-mol-payment", PostPayment).Methods("POST")
}
