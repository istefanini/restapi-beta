package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/api"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("alive true")
}

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

func main() {

	DBConection := ConectDB()
	r := mux.NewRouter()
	a := &api.API{}
	a.RegisterRoutes(r)

	r.Handle("/healthcheck", healthcheck.Handler(
		healthcheck.WithTimeout(5*time.Second),
		healthcheck.WithChecker(
			"database", healthcheck.CheckerFunc(
				func(ctx context.Context) error {
					return DBConection.PingContext(ctx)
				},
			),
		),
	))

	srv := &http.Server{
		Addr:    ":" + os.Getenv("API_PORT"),
		Handler: r,
	}
	srv.ListenAndServe()
}
