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
	specialFiled float64
}

func main() {
	e := SpecialEntity{
		specialFiled: 55.34,
		Entity: Entity{
			name:    "my house",
			version: "2.0",
			Position: Position{
				x: 233,
				y: 525,
			},
		},
	}

	e.id = "some id"
	e.name = "Bob"
	e.x = 434
	e.y = 222
	fmt.Printf("%+v\n", e.Position)
}
