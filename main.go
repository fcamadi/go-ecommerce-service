package main

import (
	"net/http"

	"github.com/cgrs/ecommerce-service-starter/middlewares"
	"github.com/cgrs/ecommerce-service-starter/server"
)

func main() {
	server.Start(
		server.CreateServer(
			"",
			middlewares.WithLogger(
				middlewares.WithJson(
					http.HandlerFunc(server.RootHandler),
				),
				nil,
			),
		),
	)
}
