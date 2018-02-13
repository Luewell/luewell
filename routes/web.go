package routes

import (
	"fmt"

	"github.com/luewell/framework/Illuminate/Http"
)

func Web(http *Http.Router) {
	http.Get("/", func(ctx *Http.Context) {
		fmt.Fprintln(ctx.ResponseWriter, "It Works!")
	})

	http.Get("/hello/world", func(ctx *Http.Context) {
		fmt.Fprintln(ctx.ResponseWriter, "Hello world!")
	})
}
