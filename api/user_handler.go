package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
)

type UserHadler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHadler {
	return &UserHadler{
		userStore: userStore,
	}
}

func (h *UserHadler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
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
