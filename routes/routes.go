package routes

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/handlers"
	"github.com/istefanini/restapi-beta/middleware"
)

func RegisteredRoutes(r *mux.Router){
	r.Use(requestIDHandler)
	r.HandleFunc("/payments", handlers.GetPayments).Methods("GET")
}