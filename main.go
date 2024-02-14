package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/api"
	"github.com/phuslu/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbuei = "mongodb://localhost:27017"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	fmt.Println(client)

	listerAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listerAddr)
}
