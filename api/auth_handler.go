package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
)

type AuthHandeler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandeler {
	return &AuthHandeler{
		userStore: userStore,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandeler) HandleAuthenticate(c *fiber.Ctx) error {
	var AuthParams AuthParams
	if err := c.BodyParser(&AuthParams); err != nil {
		return err
	}
	fmt.Println(AuthParams)
	return nil
}
