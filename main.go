package main

// 1. advanced interface mechanics
// 2. typed functions
// * these represent the composability of your program
// try to keep your interfaces as short as possible

type Putter interface {
	Put(id int, val any) error
}

// so, this Storage interface could end up implementing 20 different function methods...
// in our updateValue() function below, we only need access to one (Put)
type Storage interface {
	Get(id int) (any, error)
	// Put(id int, val any) error
	Putter
}

type simplePutter struct{}

func (s *simplePutter) Put(id int, val any) error {
	return nil
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

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

func main() {
	// s := &Server{
	// 	store: &FooStorage{},
	// }
	// updateValue(1, "foo", s.store)

	sputter := &simplePutter{}
	updateValue(1, "foo", sputter)
}
