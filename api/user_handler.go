package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
)

type UserHadler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHadler {
	return &UserHadler{
		userStore: userStore,
	}
}
func (h *UserHadler) HandlePostUser (c *fiber.Ctx) error{
	return nil
}

func (h *UserHadler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHadler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return nil
	}
	fmt.Println(users)
	return c.JSON(users)
}
