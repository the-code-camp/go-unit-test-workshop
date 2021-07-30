package domain

import "time"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth time.Time
	Status      string
}
