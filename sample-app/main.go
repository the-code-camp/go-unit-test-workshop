package main

import (
	"github.com/gorilla/mux"
	"github.com/the-code-camp/unit-test-workshop/sample-app/routes"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/customers", routes.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", routes.SaveNewCustomer).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8181", r))
}
