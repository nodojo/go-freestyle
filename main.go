package main

import "fmt"

// struct embedding makes it easier to make adjustments to the struct field members,
// which then filter down into the structs in which they are embedded...
// it also removes code duplication

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity       // struct embedding
	specialField float64
}

func main() {
	e := SpecialEntity{
		specialField: 88.88,
		Entity: Entity{
			name:    "my special entity",
			version: "1.1",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
	}

	// alternatively, you can initialize the field individually
	e.id = "my special id"
	e.name = "foo"

	// even though the struct is embedded, you can still directly access the field member
	e.x = 3333
	e.y = 1111

	fmt.Printf("%+v\n", e)
	// fmt.Printf("%+v\n", e.id)
}
