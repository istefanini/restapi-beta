package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/istefanini/restapi-beta/models"
)

type API struct{}

func ConectDB() (conection *sql.DB) {
	Driver := "sqlserver"
	Username := "serviceweb"
	Password := "Condor551"
	Host := "172.16.1.144"
	Instance := "dv"
	database := "Interoperabilidad"

	conection, err := sql.Open(Driver, Driver+"://"+Username+":"+Password+"@"+Host+"/"+Instance+"?"+"database="+database+"&"+"encrypt=disable")
	if err != nil {
		panic(err.Error())
	}
	return conection
}

func PostPayment(w http.ResponseWriter, r *http.Request) {

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var newPayment models.Payment
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newPayment)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	ctx := context.Background()
	DBConection := ConectDB()
	tsql := fmt.Sprintf("USE [Interoperabilidad] INSERT INTO [dbo].[NotificationMOLPayment]([Key],[External_Reference],[Status],[Amount]) VALUES (@Key, @External_reference, @Status, @Amount);")
	result, err2 := DBConection.ExecContext(
		ctx,
		tsql,
		sql.Named("Key", newPayment.Key),
		sql.Named("External_reference", newPayment.External_reference),
		sql.Named("Status", newPayment.Status),
		sql.Named("Amount", newPayment.Rate),
	)
	if err2 != nil {
		errorResponse(w, "Error inserting new row: "+err2.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(result)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
