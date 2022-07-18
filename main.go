package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/api"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Rest api funcionando")
}

func init() {
	gotenv.Load()
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("alive true")
}

func main() {

	log.Println(os.Getenv("x-api-key"))

	r := mux.NewRouter()

	a := &api.API{}

	a.RegisterRoutes(r)

	r.HandleFunc("/", IndexRoute).Methods(http.MethodGet)
	r.HandleFunc("/health", HealthCheckHandler)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	fmt.Println("Server corriendo en puerto", srv.Addr)
	srv.ListenAndServe()
}
