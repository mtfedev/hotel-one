package main

import "fmt"

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func main() {
	player := Player{
		name:        "Bob",
		health:      101,
		attackPower: 40.2}
	// when I added "+" I see everything in %v
	// format for int is "d"
	fmt.Printf("health: %d\n", player.health())
}
