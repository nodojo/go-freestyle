package main

import "fmt"

// global variables declared outside of functions
//var name = "foo" // direct assignment, type is automatically inferred (this is only possible in the global scope when using the keyword var)
//var firstName string = "foo" // direct assignment, type is explicitly defined
//var lastName string // initialized as the default value of the type (empty string in this case)

// preferred declaration format versus the declaration above
// var (
// 	name             = "foo"
// 	firstName string = "foo"
// 	lastName  string
// )

// constants are almost always only defined in the global namespace
// const version = 1

// constant variable names always begin with a lowercase character
const (
	version = 1
	keyLen  = 10
)

func main() {
	// local variables don't require you to use the var keyword
	// version := 1           // the := operator means the type will be inferred (int in this case)
	// otherVersion := "bar"  // inferred as a string by the compiler
	// anotherVersion := 10.1 // inferred as a float32 or float64 (float32 in this case due to the small size of the number)

	// fmt.Println(version)
	// fmt.Println(otherVersion)
	// fmt.Println(anotherVersion)

	// var version int

	fmt.Println(version)
}
