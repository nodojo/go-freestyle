package main

import "fmt"

func main() {
	// FOR LOOPS
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("it:", i)
	// }

	// numbers := []int{1, 2, 3, 4, 5, 6, 7}

	// len() returns the length of the slice that is passed in as its parameter
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// RANGES
	// names := []string{"a", "b", "c", "d"}

	// here, the variable i is the iterator -> it outputs like 0, 1, 2, 3
	// for i := range names {
	// 	fmt.Println(i)
	// }

	// use the underscore character as the iterator since it is not being referenced
	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	// BREAK
	// for _, name := range names {
	// 	if name == "a" {
	// 		break
	// 		// return
	// 	}
	// }

	// fmt.Println("exited loop")

	// MAP
	// NOTE: in go, maps are unsorted by default.. so the values it stores will always be randomized
	// users := map[string]int{
	// 	"foo":   1,
	// 	"bar":   2,
	// 	"baz":   3,
	// 	"Alice": 33,
	// 	"Bob":   88,
	// }

	// for key, value := range users {
	// 	fmt.Printf("key %s value %d\n", key, value)
	// }

	// SWITCH
	name := "foo"

	switch name {
	case "Alice":
		fmt.Println("the name is Alice")
	case "Bob":
		fmt.Println("the name is Bob")
	default:
		fmt.Println("default case triggered with the name =>", name)
	}
}
