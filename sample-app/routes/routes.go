package routes

import (
	"encoding/json"
	"github.com/the-code-camp/unit-test-workshop/sample-app/dto"
	"github.com/the-code-camp/unit-test-workshop/sample-app/service"
	"log"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	allCustomers, err := service.CustomerService.GetAllCustomers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		err := json.NewEncoder(w).Encode(allCustomers)
		if err != nil {
			log.Println("not able to send response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func SaveNewCustomer(w http.ResponseWriter, r *http.Request) {
	var c dto.CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println("Invalid request" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCustomerId, err := service.CustomerService.SaveCustomer(c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write([]byte(newCustomerId)); err != nil {
			log.Println("not able to send response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
