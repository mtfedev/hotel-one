package main

import (
	"fmt"

	"github.com/mtfedev/hotel-one/types"
)

func main() {
	hotel := types.Hotel{
		Name:     "Ballucia",
		Location: "Paris",
	}

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrise: 88.9,
	}

	fmt.Println("seeding the database")
}
