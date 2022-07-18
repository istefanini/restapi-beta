package api

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/models"
)

type API struct{}

var PaymentsData = models.Payments{
	{
		Key:                "h_di!asjcoui2wsc22hcjdaiaaaaiu8x.asi+cn-na",
		External_reference: 10012151,
		Rate:               1050.50,
		Status:             "Completed",
	},
	{
		Key:                "h_di!asjcoui2wsc22hcjdaiaaaaiu8x.asi+cn-nb",
		External_reference: 10012152,
		Rate:               370,
		Status:             "Pendant",
	},
	{
		Key:                "h_di!asjcoui2wsc22hcjdaiaaaaiu8x.asi+cn-nz",
		External_reference: 10012153,
		Rate:               250.80,
		Status:             "Completed",
	},
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaymentsData)
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var newPayment models.Payment

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error procesando el pago")
	}

	json.Unmarshal(reqBody, &newPayment)
	PaymentsData = append(PaymentsData, newPayment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPayment)

}

func GetOnePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paymentID, err := strconv.Atoi(vars["external_reference"])
	if err != nil {
		fmt.Fprintf(w, "Invalid external_reference")
		return
	}

	for _, payment := range PaymentsData {
		if payment.External_reference == paymentID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(payment)
		}
	}
}