package main

// 1. advanced interface mechanics
// 2. typed functions
// * these represent the composability of your program
// try to keep your interfaces as short as possible

// so, this Storage interface could end up implementing 20 different function methods...
// in our updateValue() function below, we only need access to one (Put)
type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type BarStorage struct{}

func (s *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *BarStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

// now we use dependency injection to remove the dependency
// this would be the way we could cleanly "plug & play" storage methods...
// (ex: Postgres, MongoDB, testing (which is an "in-memory" store, so no need to connect over the network))
func updateValue(id int, val any, store Storage) error {
	return store.Put(id, val)
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	updateValue(1, "foo", s.store)
}
