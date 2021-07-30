package service

import (
	"errors"
	"github.com/the-code-camp/unit-test-workshop/sample-app/domain"
	"github.com/the-code-camp/unit-test-workshop/sample-app/dto"
	"log"
	"time"
)

type Service struct{}

var (
	CustomerService Service = Service{}
)

func (s Service) GetAllCustomers() ([]dto.CustomerResponse, error) {
	all, err := domain.CustomerRepo.FindAll()
	if err != nil {
		return nil, err
	}
	customers := make([]dto.CustomerResponse, 0)
	for _, c := range all {
		customer := dto.CustomerResponse{
			Id:          c.Id,
			Name:        c.Name,
			City:        c.City,
			Zipcode:     c.Zipcode,
			DateofBirth: c.DateofBirth.Format("2006-01-02"),
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (s Service) SaveCustomer(c dto.CustomerRequest) (string, error) {
	if c.Name == "" || c.City == "" || c.Zipcode == "" || c.DateofBirth == "" {
		return "", errors.New("name, city and zipcode cannot be empty")
	}

	parsedDate, err := time.Parse("2006-01-02", c.DateofBirth)
	if err != nil {
		log.Println("Incorrent date format" + err.Error())
		return "", errors.New("incorrect date format. Should be YYYY-MM-DD")
	}
	customer := domain.Customer{
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: parsedDate,
	}

	insertId, err := domain.CustomerRepo.Save(customer)
	if err != nil {
		return "", err
	}
	return insertId, nil
}
