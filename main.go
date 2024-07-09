package main

import "fmt"

type Entity struct {
	name    string //
	id      string //
	version string //--> field members
	posx    int    //
	posy    int    //
}

func main() {
	// e := &Entity{ // declared as pointer
	e := Entity{ // declared as literal
		name:    "my entity",
		id:      "id 1",
		version: "version 1.1",
		posx:    100,
		posy:    200,
	}
	fmt.Printf("%+v\n", e)
}
