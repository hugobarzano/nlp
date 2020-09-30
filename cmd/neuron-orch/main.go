
package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"neuron-orch/internal/service"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", service.Ping).Methods(http.MethodGet)
	router.HandleFunc("/msg", service.Receiver).Methods(http.MethodPost)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Println("Listening on http://0.0.0.0:8080 !")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}