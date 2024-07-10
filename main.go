package main

// we're gonna make this generic
type CustomMap struct {
	data map[string]int // example use: usernames with an id or age
}

// implement a function into this map to insert values
func (m *CustomMap) Insert(k string, v int) error {
	// the problem is we have a big dependency... our map, because it's a string/int -> if we want to change this, we can't
	// in other words, we're tied to the data types used by the map
	// how are we going to fix this? generics
	m.data[k] = v
	return nil
}

// make a new map to initialize the data
// it's going to be a pointer to this CustomMap and it's going to return the pointer to the CustomMap
func NewCustomMap() *CustomMap {
	// the data inside of this map is going to be make me a map of string to int
	return &CustomMap{
		data: make(map[string]int),
	}
}

func main() {

}
