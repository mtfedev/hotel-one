package api

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
	"go.mongodb.org/mongo-driver/mongo"
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

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func (h *AuthHandeler) HandleAuthenticate(c *fiber.Ctx) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("invalid credentials")
		}
		return err
	}

	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		return fmt.Errorf("invalid credentials")

	}
	resp := AuthResponse{
		User:  user,
		Token: createTotenFromUser(user),
	}
	return c.JSON(resp)
}

func createTotenFromUser(user *types.User) string {
	now := time.Now()
	expiers := now.Add(time.Hour * 4).Unix()
	claims := jwt.MapClaims{
		"id":      user.ID,
		"emial":   user.Email,
		"expiers": expiers,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sing to token with secert", err)
	}
	return tokenStr
}
