package main

import (
	"sync"
)

// state -> set OR update OR delete (how we manipulate state)

type State struct {
	mu    sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count = i
}

func main() {
}
