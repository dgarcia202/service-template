package main

import "github.com/dgarcia202/service-template/pkg/app"

func main() {
	app := app.New()
	app.ServiceName = "customers"
	app.ShortDescription = "Hugo is a very fast static site generator"

	app.LongDescription = `A Fast and Flexible Static Site Generator built with
	love by spf13 and friends in Go.
	Complete documentation is available at http://hugo.spf13.com`

	app.Version = "0.0.1"

	app.Run()
}
