package middleware

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWT auth")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if ok {
		return fmt.Errorf("unauthorized")
	}

	fmt.Println("token", token)
	return nil
}

func parseJWTToken(tokenStr string) error {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Println(claims.Foo, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

}
