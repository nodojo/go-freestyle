package main

// 1. advanced interface mechanics
// 2. typed functions
// * these represent the composability of your program

type Storage interface {
	// this is going to be get by id
	// Get(int) (any, error)
	// you don't have to specify id here, but you can and it will be clearer and more verbose
	Get(id int) (any, error)
	Put(id int, val any) error
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{} // instantiate
	s.store.Get(1)
	s.store.Put(1, "foo")
	// if you run this, you'll get a runtime error because the Storage interface is empty
}
