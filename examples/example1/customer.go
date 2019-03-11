package main

import (
	"net/http"

	app "github.com/dgarcia202/service-template/pkg/app"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getAllCustomers(c *gin.Context) {
	log.Trace("Querying all customers")
	var customers []Customer
	app.Db().Preload("Addresses").Find(&customers)
	c.JSON(http.StatusOK, customers)
}

func getCustomerByID(c *gin.Context) {
	id := c.Param("id")
	log.Trace("Querying customer with id ", id)
	var customer Customer
	if app.Db().Where("ID = ?", id).Preload("Addresses").First(&customer).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func addCustomer(c *gin.Context) {
	log.Trace("Adding a new customer")
	var json customerDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	customer := Customer{Name: json.Name, LegalName: json.LegalName, Addresses: make([]Address, 0, 1)}
	app.Db().Create(&customer)
	c.JSON(http.StatusCreated, customer)
}

func updateCustomer(c *gin.Context) {
	id := c.Param("id")
	log.Trace("Updating customer with id ", id)
	var json customerDto
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var customer Customer
	if app.Db().Where("ID = ?", id).First(&customer).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	customer.Name = json.Name
	customer.LegalName = json.LegalName
	app.Db().Save(&customer)

	c.JSON(http.StatusOK, customer)
}

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")
	log.Trace("Removing customer with id ", id)
	var customer Customer
	if app.Db().Where("ID = ?", id).First(&customer).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	app.Db().Delete(&customer)
	c.Status(http.StatusOK)
}

func defineCustomerRoutes(r *gin.Engine) {
	r.GET("/customers", getAllCustomers)
	r.GET("/customers/:id", getCustomerByID)
	r.POST("/customers", addCustomer)
	r.PUT("/customers/:id", updateCustomer)
	r.DELETE("/customers/:id", deleteCustomer)
}
