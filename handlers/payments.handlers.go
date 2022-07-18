package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/data"
	"github.com/istefanini/restapi-beta/models"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var newPayment models.Payment

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error procesando el pago")
	}

	json.Unmarshal(reqBody, &newPayment)
	data.PaymentsData = append(data.PaymentsData, newPayment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPayment)

}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.PaymentsData)
}

func GetOnePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["external_reference"])
	if err != nil {
		fmt.Fprintf(w, "Invalid external_reference")
		return
	}

	for _, payment := range data.PaymentsData {
		if payment.External_reference == paymentID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(payment)
		}
	}
}
