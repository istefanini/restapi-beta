package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/istefanini/restapi-beta/handlers"
)

func main() {
	r := mux.NewRouter()

	a := 

	a.RegisterRoutes(r)

	
	r.HandleFunc("/", handlers.IndexRoute).Methods(http.MethodGet)

	srv:= &http.Server{
		Addr: ":3000",
		Handler: r,
	}

	fmt.Println("Server corriendo en puerto", srv.Addr)
	srv.ListenAndServe()
}
