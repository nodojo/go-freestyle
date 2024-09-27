package main

import (
	"sync"
)

// state -> set OR update OR delete (how we manipulate state)

type State struct {
	sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.count = i
}

func main() {
}
