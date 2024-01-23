package main

import "fmt"

// if player is not a pointer we are adjusting the copu of the player
// not the actual player
type Player struct {
	HP int
	// data
}

// function receiver
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage, New Hp ->", p.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}
	player.TakeDamage(10)

	fmt.Printf("%+v\n", player)
}
