package main

import (
	"net/http"

	app "github.com/dgarcia202/service-template/pkg/app"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func addAddressToCustomer(c *gin.Context) {
	customerID := c.Param("id")
	log.Trace("Adding address to customer ", customerID)
	var customer Customer
	if app.Db().Where("ID = ?", customerID).Preload("Addresses").First(&customer).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var json addressDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	address := Address{
		CustomerID:   customerID,
		AddressLine1: json.AddressLine1,
		AddressLine2: json.AddressLine2,
		City:         json.City,
		State:        json.State,
		ZipCode:      json.ZipCode,
		Country:      json.Country}

	customer.Addresses = append(customer.Addresses, address)
	app.Db().Save(&customer)

	c.JSON(http.StatusOK, customer)
}

func getAllCustomerAddresses(c *gin.Context) {
	customerID := c.Param("id")
	log.Trace("Querying all addresses from customer ", customerID)

	var count int
	app.Db().Model(&Customer{}).Where("ID = ?", customerID).Count(&count)
	if count == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var addresses []Address
	app.Db().Where("CUSTOMER_ID = ?", customerID).Find(&addresses)
	c.JSON(http.StatusOK, addresses)
}

func defineAddressRoutes(r *gin.Engine) {
	r.POST("/customers/:id/addresses", addAddressToCustomer)
	r.GET("/customers/:id/addresses", getAllCustomerAddresses)
}
