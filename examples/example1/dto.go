package main

type customerDto struct {
	Name      string `json:"name"`
	LegalName string `json:"legalName"`
}

type addressDto struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
}
