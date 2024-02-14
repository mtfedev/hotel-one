package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/api"
)

func main() {
	listerAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listerAddr)
}
