package main

import "fmt"

type Color int

// implement the stringer interface (basically a fmt.Stringer)
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "PINK"
	default:
		panic("invalid color given")
	}
}

const (
	ColorBlue Color = iota // remember, iota means increment everything below (starting with zero)... use iota + 1 to start with 1
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	// if a stringer interface has been implemented, the go compiler will first attempt to use that
	fmt.Println("the color is:", ColorBlack)
}
