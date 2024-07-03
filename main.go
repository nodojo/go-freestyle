package main

import (
	"fmt"
)

// #6 interfaces

// Storage
// Stor(er) // idiomatic way of naming interfaces in go
// io.Read(er)
// io.Writ(er)

// declaration (not implementation obv)
type NumberStorer interface {
	// describe here what a storer can do
	// going to retrieve everything from the db
	GetAll() ([]int, error)
	// in an interface, we don't need to specify the exact variable name,
	// because we're just describing what it can do... we're not doing
	// an implementation yet, so it doesn't really matter
	// Put(number int) error // so this is unnecessary
	Put(int) error // this is more idiomatic
}

// to implement the NumberStorer interface all we need to do is implement its two functions
type MongoDBNumberStore struct {
	// some values
}

// remember we learned previously that we can attach functions to our structures -> called function receivers
func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into teh mongoDB storage")
	return nil
}

// now, since we've used this interface pattern, if there is a sudden need to pivot to a different
// type of data store (here we're pivoting from mongodb to postgres), it's not a problem
type PostgresNumberStore struct {
	// postgres values (db connection)
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into teh postgres storage")
	return nil
}

type ApiServer struct {
	// lowercase numberStore means it's private, uppercase NumberStore would be public? (I think)
	numberStore NumberStorer // loosely coupled -> preferred
	// you see people doing things like the following, which is highly coupled...
	// what happens if we decide to change to an alternate form of storage?
	// so, this is a bad behavior.. it's hard to test, hard to test, hard to maintain
	// numberStore MongoDBStore
}

func main() {
	apiServer := ApiServer{
		// numberStore: MongoDBNumberStore{},
		// since we've used this interface pattern, it's simple for us to update the data store
		numberStore: PostgresNumberStore{},
	}

	// logic below
	// err := apiServer.numberStore.Put(1)
	// if err != nil {
	// 	panic(err)
	// }
	// shorthand version of block above
	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	// if no error
	fmt.Println(numbers)
}
