package main

import "fmt"

// we're gonna make this generic
// in the brackets you're going to provide your generic types
// our key (K) can be anything, but we need to say that's it's going to be a comparable...
// ex: key in a hash map that needs to be compared to another key in a hash map to see if it exists
type CustomMap[K comparable, V any] struct {
	data map[K]V // example use: usernames with an id or age
}

// implement a function into this map to insert values
func (m *CustomMap[K, V]) Insert(k K, v V) error {
	// the problem is we have a big dependency... our map, because it's a string/int -> if we want to change this, we can't
	// in other words, we're tied to the data types used by the map
	// how are we going to fix this? generics
	m.data[k] = v
	return nil
}

// make a new map to initialize the data
// it's going to be a pointer to this CustomMap and it's going to return the pointer to the CustomMap
func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	// the data inside of this map is going to be make me a map of string to int
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	// notice we initialize the map using the data types we want
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)
	fmt.Printf("map1: %+v\n", m1)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 100.3333)
	fmt.Printf("map2: %+v\n", m2)
}
