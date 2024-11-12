package api

import "github.com/gofiber/fiber/v2"

// be aware, we are in a different package -> we need to make
// our functions, variables, structures, types etc. public (uppercase)
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
