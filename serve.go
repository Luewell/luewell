package main

import (
	"os"

	"github.com/luewell/framework/Illuminate/Http"
	"github.com/luewell/framework/bootstrap"
	"github.com/luewell/luewell/routes"
)

func init() {
	bootstrap.App()
}

func main() {

	/* RouteServiceProvider */
	http := Http.Init()

	routes.Web(http)

	http.Listen(os.Getenv("APP_URL"))
}
