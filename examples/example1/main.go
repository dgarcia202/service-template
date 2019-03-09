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

func setupRoutes(r *gin.Engine) {
	r.GET("/customers", func(c *gin.Context) {
		app.Db().Create(&Customer{Name: "Acme LTD.", LegalName: "TEXAS ACME INC. LTD."})
		c.String(http.StatusOK, "adios")
	})
}

func main() {
	app.ServiceName("customers")
	app.ShortDescription("This is a dummy customers service")
	app.LongDescription(`Customers service is an example micro service developed just
		for educational purposes`)

	app.Version("0.0.1")
	app.AddModel(&Customer{})
	app.AddHTTPSetup(setupRoutes)
	app.Run()
}
