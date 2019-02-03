package main

import "github.com/dgarcia202/service-template/pkg/app"

func main() {
	app := app.New()
	app.ServiceName = "customers"
	app.ShortDescription = "This is a dummy customers service"

	app.LongDescription = `Customers service is an example micro service developed just
		for educational purposes`

	app.Version = "0.0.1"

	app.Run()
}
