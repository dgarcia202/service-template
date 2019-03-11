package main

import (
	app "github.com/dgarcia202/service-template/pkg/app"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app.ServiceName("customers")
	app.ShortDescription("This is a dummy customers service")
	app.LongDescription(`Customers service is an example micro service developed just
		for educational purposes`)

	app.Version("0.0.1")
	app.AddModels(&Customer{}, &Address{})
	app.AddHTTPSetup(defineCustomerRoutes)
	app.AddHTTPSetup(defineAddressRoutes)
	app.Run()
}
