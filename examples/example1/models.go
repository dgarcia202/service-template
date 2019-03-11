package main

import (
	app "github.com/dgarcia202/service-template/pkg/app"
)

// Customer entity
type Customer struct {
	app.Model
	Name      string
	LegalName string
	Addresses []Address
}

// Address of a customer
type Address struct {
	app.Model
	CustomerID   string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	ZipCode      string
	Country      string
}
