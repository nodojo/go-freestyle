package main

import "fmt"

// state -> set OR update OR delete (how we manipulate state)

type State struct {
	count int
}

func main() {
	// initialize state (dont forget if the fields arent set, go will initialize them to the default values for their type)
	state := State{}

	for i := 0; i < 10; i++ {
		state.count = i
	}

	fmt.Println(state)
}
