package main

import (
	"fmt"
)

// a pointer is an 8 byte long integer that points to a memory address
type Player struct {
	HP int
}

// function receiver
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New HP ->", p.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}

	// send pointer to player location in memory
	player.TakeDamage(10)

	fmt.Printf("%+v\n", player)
}
