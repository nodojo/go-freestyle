package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife
type WeaponType int

// String() functions can be "attached" to custom types (our enum in this case) so that a string value can be set and accessed
func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenStick:
		return "WOODENSTICK"
	case Knife:
		return "KNIFE"
	}

	return ""
}

// // manually typing in incrementing integers
// const (
// 	Axe         WeaponType = 1
// 	Sword       WeaponType = 2
// 	WoodenStick WeaponType = 3
// 	Knife       WeaponType = 4
// )

// using iota which sets the first property value to 0 then increments the value set for each following property
const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		// panic is an exception that stops the program and outputs a stack trace with the message you set for it
		panic("weapon does not exist")
	}
}

// // first example (uses strings instead of our "enums")
// func getDamage(weaponType string) int {
// 	switch weaponType {
// 	case "axe":
// 		return 100
// 	case "sword":
// 		return 90
// 	case "woodenStick":
// 		return 1
// 	case "knife":
// 		return 40
// 	default:
// 		// panic is an exception that stops the program and outputs a stack trace with the message you set for it
// 		panic("weapon does not exist")
// 	}
// }

func main() {
	// fmt.Println("weapon damage: ", getDamage("knife"))
	// fmt.Println("weapon damage: ", getDamage("nife")) // test panic case
	// fmt.Println("weapon damage: ", getDamage(Knife))

	// by using %s here, we're telling the compiler to detect if the %s value (Axe, Sword, etc. below) has a String() function attached, then it should output the corresponding string value
	fmt.Printf("weapon damage (%s) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("weapon damage (%s) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("weapon damage (%s) (%d):\n", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("weapon damage (%s) (%d):\n", Knife, getDamage(Knife))
}
