package api

import (
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
func (h *UserHadler) HandlePutUser(c *fiber.Ctx) error {
	return nil
}
func (h *UserHadler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": userID})
}

func (h *UserHadler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
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
		mongo.
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
