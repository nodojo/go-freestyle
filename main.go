package main

import "fmt"

// a pointer is an 8 byte long integer that points to a memory address
type Player struct {
	HP int
}

func TakeDamage(player Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}
	// golang sends a copy of player as the param to TakeDamage()...
	// which means, if player is not a pointer then we are adjusting the copy of the player, not the actual player
	TakeDamage(player, 10)

	fmt.Printf("%+v\n", player)
}
