package main

import "fmt"

type Entity struct {
	name    string //
	id      string //
	version string //--> field members
	posx    int    //
	posy    int    //
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
			// id:      "my special id",
			posx: 100,
			posy: 200,
		},
	}

	// alternatively, you can initialize the field individually
	e.id = "my special id"

	fmt.Printf("%+v\n", e)
	// fmt.Printf("%+v\n", e.id)
}
