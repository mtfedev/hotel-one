package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		fmt.Println("token not ptesnt in the header")
		return fmt.Errorf("unauthorized")
	}

	claims, err := validateToken(token)
	if err != nil {
		return err
	}
	expires := claims["expires"].(time.Time)
	if time.Now().After(expires) {
		return fmt.Errorf("token expired")
	}
	fmt.Println(expires)
	return nil
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("ivalet single method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SICRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT tocken", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		fmt.Errorf("ivalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
