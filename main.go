package main

import "fmt"

// a pointer is an 8 byte long integer that points to a memory address
type Player struct {
	HP int
}

// accept pointer to player instead of passing a copy by value
func TakeDamage(player *Player, amount int) {
	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}

	// send pointer to player location in memory
	TakeDamage(&player, 10)

	fmt.Printf("%+v\n", player)
}
