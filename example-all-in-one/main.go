package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers", saveNewCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8181", r))
}

type CustomerRequest struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
}

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

func saveNewCustomer(w http.ResponseWriter, r *http.Request) {
	var c CustomerRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Println("Invalid request" + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if c.Name == "" || c.City == "" || c.Zipcode == "" || c.DateofBirth == "" {
		handleError("name, city and zipcode cannot be empty", w, http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", c.DateofBirth)
	if err != nil {
		log.Println("Incorrent date format" + err.Error())
		handleError("incorrect date format. Should be YYYY-MM-DD", w, http.StatusBadRequest)
		return
	}

	sqlInsert := "INSERT INTO customers (name, date_of_birth, city, zipcode ) values (?, ?, ?, ?)"
	client := getDbClient()
	result, err := client.Exec(sqlInsert, c.Name, parsedDate, c.City, c.Zipcode)
	if err != nil {
		log.Println("Error while saving data in database" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error while fetching last inserted Id" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.FormatInt(lastInsertId, 10)))
}

func handleError(message string, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	client := getDbClient()
	customers := make([]Customer, 0)
	rows, err := client.Query("select customer_id, name, city, zipcode, date_of_birth, status from customers")
	if err != nil {
		log.Println("error while fetching from database", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		c := new(Customer)
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("error while scanning from rows", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		customers = append(customers, *c)
	}

	err = json.NewEncoder(w).Encode(map[string][]Customer{
		"customers": customers,
	})
	if err != nil {
		log.Println("error while decoding to JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getDbClient() *sql.DB {
	dbUser := "root"
	dbPasswd := "codecamp"
	dbAddr := "localhost"
	dbPort := "3306"
	dbName := "banking"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	return client
}
