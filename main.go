package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/nodojo/go-freestyle/api"
)

// run to change the port to :7000:
//
//	./bin/api --listenAddr :7000
func main() {
	listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUser)
	app.Listen(*listenAddr)
}
