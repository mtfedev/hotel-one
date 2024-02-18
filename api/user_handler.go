package api

import (
	"go/types"

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
func (h *UserHadler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
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
	return c.JSON(users)
}
