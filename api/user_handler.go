package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
)

type UserHadler struct {
	userStore db.UserStore
}

func (h *UserHadler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(id)
	id err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHadler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "at the watercooler",
	}
	return c.JSON(u)
}
