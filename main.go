package main

import "fmt"

type Position struct {
	x, y int
}
type Entity struct {
	name    string
	id      string
	version string

	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := SpecialEntity{
		specialField: 86.22,
		Entity: Entity{
			name:    "my special entity",
			version: "1.1",
			Position: Position{
				x: 23,
				y: 42,
			},
		},
	}
	e.id = "my spcial id"
	e.x = 234
	e.y = 2414
	fmt.Printf("%+v\n", e.id)
}
