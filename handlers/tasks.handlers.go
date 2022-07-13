package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/faztweb/golang-restapi-crud/data"
	"github.com/faztweb/golang-restapi-crud/models"
	"github.com/gorilla/mux"
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

// func UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	var updatedTask models.Task

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 	}

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Please Enter Valid Data")
// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, t := range data.TasksData {
// 		if t.ID == taskID {
// 			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)

// 			updatedTask.ID = t.ID
// 			data.TasksData = append(data.TasksData, updatedTask)

// 			// w.Header().Set("Content-Type", "application/json")
// 			// json.NewEncoder(w).Encode(updatedTask)
// 			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
// 		}
// 	}

// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid User ID")
// 		return
// 	}

// 	for i, t := range data.TasksData {
// 		if t.ID == taskID {
// 			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
// 			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
// 		}
// 	}
// }
