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
	s.count = i
	// note: every Lock() must have a Unlock() to free up the locked resources
	s.mu.Unlock()
}

func main() {
}
