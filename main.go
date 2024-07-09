package main

// 1. advanced interface mechanics
// 2. typed functions
// * these represent the composability of your program
// try to keep your interfaces as short as possible

type Storage interface {
	// this is going to be get by id
	// Get(int) (any, error)
	// you don't have to specify id here, but you can and it will be clearer and more verbose
	Get(id int) (any, error)
	Put(id int, val any) error
}

type FooStorage struct{}

// s is a pointer to FooStorage
func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any) error {
	// problem here is the dependency we have within our function -> this reference to FooStorage...
	// the way it's written here, we are always going to create FooStorage with this implementation
	store := &FooStorage{}
	return store.Put(id, val)
}

func main() {
	s := &Server{
		store: &FooStorage{},
	} // instantiate
	s.store.Get(1)
	s.store.Put(1, "foo")
	// if you run this, you'll get a runtime error because the Storage interface is empty...
	// that's because Storage is the interface and not the actual implementation
}
