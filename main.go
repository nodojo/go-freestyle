package main

import "fmt"

// // GENERAL TYPES
// var (
// 	floatVar32 float32 = 0.1
// 	floatVar64 float64 = 0.1
// 	name       string  = "foo"
// 	intVar32   int32   = 1
// 	intVar64   int64   = 48484
// 	intVar     int     = -1
// 	uintVar    uint    = 1   // unsigned integer, cannot be negative
// 	uint32Var  uint32  = 1   // 4 bytes
// 	uint64Var  uint64  = 1   // 8 bytes
// 	unit8Var   uint8   = 0x1 // 1 byte
// 	byteVar    byte    = 0x2 // exactly the same as a unit8
// 	runeVar    rune    = 'a' // specified with single quotes (so that it's a rune and not a string), if you index a string you're going to also have a rune
// )

// // STRUCTS (an empty struct doesn't occupy any memory)
// // golang initializes empty variables to the default state of that type of variable
// type Player struct {
// 	name        string
// 	health      int
// 	attackPower float64
// }

// // when you attach a function to a struct it's called a function receiver
// func (player Player) getHealth() int {
// 	return player.health
// }

// first struct example
// notice param format is variable name the its type in golang
// func getHealth(player Player) int {
// 	return player.health
// }

// // CUSTOM TYPE
// // telling compiler that I have a type that is called weapon and it's giong to be a string
// type Weapon string

// func getWeapon(weapon Weapon) string {
// 	// cast weapon to string to return (which is type Weapon (not string... even though the underlying type is string))
// 	return string(weapon)
// }

func main() {
	// ARRAYS
	// the size of an array remains static (unlike a slice's size which is dynamic)
	// numbers := [2]int{}
	// numbers[0] = 1
	// numbers[1] = 2

	// SLICES
	// an empty slice like this is almost always how you declare and use a slice
	numbers := []int{} // empty
	// numbers := []int{1, 2, 3}
	// when using the make keyword, we need to specify the initial size (2 in this example) of the slice (slices are dynamic, so they can expand/contract)
	// almost never need to use this
	// otherNumbers := make([]int, 2)
	otherNumbers := make([]int, 0) // empty

	// add to a slice
	// numbers = append(numbers, 1)
	// numbers = append(numbers, 10)

	// looping over a slice would be done using the same for/range loop that we used for maps

	fmt.Println(numbers)
	fmt.Println(otherNumbers)

	// MAPS
	// majority of the time, init a map like this (possibly the most idiomatic)
	// users := map[string]int{} //empty map
	// users := map[string]int{
	// 	"anothony": 36,
	// }
	// // the make built-in function makes a new default for certain types in golang
	// users := make(map[string]int) // hey go compiler, make me a map of strings to integers
	// users["foo"] = 10
	// users["bar"] = 11
	// delete by key from map... only applicable to maps
	// delete(users, "foo")

	// // range over a map (key and value)
	// for k, v := range users {
	// 	// %s = string, %d = integers
	// 	fmt.Printf("the key %s and the value %d\n", k, v)
	// }

	// // range over a map (keys only)
	// for key := range users {
	// 	fmt.Printf("the key %s ", key)
	// }

	// age := users["foo"]
	// since "baz" doesn't exist, it will output the default value of an integer (0)
	// this built-in go behvior can come back to shoot you in the foot, because 0 could be misinterpreted as value that has been set (not by default)
	// age := users["baz"]
	// this prevents the built-in behavior that defaults an integer to 0 from shooting you in the foot
	// age, ok := users["baz"]
	// if !ok {
	// 	fmt.Println("baz doesn't exist in the map")
	// } else {
	// 	fmt.Println("baz exists in the map: ", age)
	// }
	// fmt.Printf("age: %+v\n", age)

	// fmt.Printf("users: %+v\n", users)

	// STRUCTS
	// player := Player{
	// 	name:        "captain jack",
	// 	health:      101,
	// 	attackPower: 45.1,
	// }

	// print formatted
	// %v = verbose, %+v outputs property names in struct along with the values ("super" verbose mode)
	// fmt.Printf("this is the player: %+v\n", player)
	// %d is for integers
	// fmt.Printf("health: %d\n", getHealth(player))
	// fmt.Printf("health: %d\n", player.getHealth())
	// getHealth() was just to create an example, you don't need a getter because you can just access the same value with player.health
	// fmt.Printf("health: %d\n", player.health)
}
