package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Ballucia",
		Location: "Paris",
	}
	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrise: 88.9,
		}, {
			Type:      types.SingleRoomType,
			BasePrise: 88.9,
		}, {
			Type:      types.SeaSideRoomType,
			BasePrise: 123.9,
		},
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedRoom)
	}
}
