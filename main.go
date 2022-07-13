package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/faztweb/golang-restapi-crud/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.IndexRoute)
	router.HandleFunc("/payment/api/v1/notificaction-mol-payment ", handlers.CreatePayment).Methods("POST")
	router.HandleFunc("/payment/api/v1/payments", handlers.GetPayments).Methods("GET")
	router.HandleFunc("/payment/api/v1/payments/{external_reference}", handlers.GetOnePayment).Methods("GET")
	fmt.Println("Server corriendo en puerto", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))
}
