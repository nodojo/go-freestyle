package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nodojo/go-freestyle/types"
)

// be aware, we are in a different package -> we need to make
// our functions, variables, structures, types etc. public (uppercase)
func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the water cooler",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
