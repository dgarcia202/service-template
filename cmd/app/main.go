package main

import (
	"net/http"

	"github.com/dgarcia202/service-template/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/hola", func(c *gin.Context) {
		c.String(http.StatusOK, "adios")
	})
}

func main() {
	app := app.Instance()
	app.ServiceName = "customers"
	app.ShortDescription = "This is a dummy customers service"

	app.LongDescription = `Customers service is an example micro service developed just
		for educational purposes`

	app.Version = "0.0.1"

	app.SetupRoutes(setupRoutes)

	app.SetupRoutes(func(r *gin.Engine) {
		r.GET("/second", func(c *gin.Context) {
			c.String(http.StatusOK, viper.GetString("loglevel"))
		})
	})

	app.Run()
}
