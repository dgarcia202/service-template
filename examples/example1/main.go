package main

import (
	"net/http"

	app "github.com/dgarcia202/service-template/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Customer entity
type Customer struct {
	gorm.Model
	Name      string
	LegalName string
}

type customerDto struct {
	Name      string `json:"name"`
	LegalName string `json:"legalName"`
}

func defineRoutes(r *gin.Engine) {
	r.GET("/customers", func(c *gin.Context) {
		var customers []Customer
		app.Db().Find(&customers)
		c.JSON(http.StatusOK, customers)
	})

	r.GET("/customers/:id", func(c *gin.Context) {
		var customer Customer
		if app.Db().Where("ID = ?", c.Param("id")).First(&customer).RecordNotFound() {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, customer)
	})

	r.POST("/customers", func(c *gin.Context) {
		var json customerDto
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		customer := Customer{Name: json.Name, LegalName: json.LegalName}
		app.Db().Create(&customer)
		c.JSON(http.StatusCreated, customer)
	})

	r.PUT("/customers/:id", func(c *gin.Context) {
		var json customerDto
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		var customer Customer
		if app.Db().Where("ID = ?", c.Param("id")).First(&customer).RecordNotFound() {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		customer.Name = json.Name
		customer.LegalName = json.LegalName
		app.Db().Save(&customer)

		c.JSON(http.StatusOK, customer)
	})

	r.DELETE("/customers/:id", func(c *gin.Context) {
		var customer Customer
		if app.Db().Where("ID = ?", c.Param("id")).First(&customer).RecordNotFound() {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		app.Db().Delete(&customer)
		c.Status(http.StatusOK)
	})
}

func main() {
	app.ServiceName("customers")
	app.ShortDescription("This is a dummy customers service")
	app.LongDescription(`Customers service is an example micro service developed just
		for educational purposes`)

	app.Version("0.0.1")
	app.AddModel(&Customer{})
	app.AddHTTPSetup(defineRoutes)
	app.Run()
}
