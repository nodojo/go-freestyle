package main

import (
	"fmt"

	"github.com/nodojo/go-freestyle/types"
	"github.com/nodojo/go-freestyle/util"
)

// the two different things you can do with golang:
// --program--
// 1. something that runs / executes code (go run somefile.go)
// => will have a main() function -> entry function of thr program
// => this is going to be package main
// --package / library / module--
// 2. something that someone can import into their program or library
// => this is going to be package <whatever>

// User => public access everywher (because it begins with a capital letter)
// user => private acess BUT public in its own package
func main() {
	// number := getNumber()
	// fmt.Println("the number is:", number)

	// before running this, we need to rebuild again because we'ev added the User struct since our last build
	// now that the User struct declaration is in the types directory we just created, we are getting an undeclared name error
	// note: everything that beings with a capital letter is publically accessible, everything begining with a lowercase letter is private
	user := types.User{
		Username: util.GetUsername(),
		Age:      util.GetAge(),
	}
	// by using %+v here with Printf(), we tell the compiler to output the struct property names along with their values
	fmt.Printf("the user is %+v:", user)
}
