package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "at the watercooler",
		return c.JSON(u)
	}
	return c.JSON("James")
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
