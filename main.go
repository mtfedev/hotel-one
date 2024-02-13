package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listerAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	app.Get("/foo", handelFoo)
	apiv1.Get("/user")
	app.Listen(*listerAddr)
}

func handelFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}
