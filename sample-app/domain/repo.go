package domain

import (
	"errors"
	"github.com/the-code-camp/unit-test-workshop/sample-app/datasource"
	"log"
	"strconv"
)

type Repository struct {
}

var (
	CustomerRepo Repository = Repository{}
)

func (r Repository) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)
	rows, err := datasource.Client.Query("select customer_id, name, city, zipcode, date_of_birth, status from customers")
	if err != nil {
		log.Println("error while fetching from database", err.Error())
		return nil, errors.New("error while fetching from database")
	}

	for rows.Next() {
		c := new(Customer)
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("error while scanning from rows", err.Error())
			return nil, errors.New("error while scanning customers")
		}
		customers = append(customers, *c)
	}
	return customers, nil
}

func (r Repository) Save(c Customer) (string, error) {
	sqlInsert := "INSERT INTO customers (name, date_of_birth, city, zipcode ) values (?, ?, ?, ?)"
	result, err := datasource.Client.Exec(sqlInsert, c.Name, c.DateofBirth, c.City, c.Zipcode)
	if err != nil {
		log.Println("Error while saving data in database" + err.Error())
		return "", errors.New("error while saving customer to database")
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error while fetching last inserted Id" + err.Error())
		return "", errors.New("error while fetching last inserted id")
	}
	return strconv.FormatInt(lastInsertId, 10), nil
}
