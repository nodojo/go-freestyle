package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

// run to change the port to :7000:
//
//	./bin/api --listenAddr :7000
func main() {
	listenAddr := flag.String("listenAddr", ":5000", "the listen address of the api server")
	flag.Parse()

	app := fiber.New()
	// apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	// appv1.Get("/user", handleUser)
	app.Get("/user", handleUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"mgs": "working just fine!"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Foo"})
}
