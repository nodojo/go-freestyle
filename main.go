package main

import "sync/atomic"

// atomic values are another way of working with things in a synchronized way

type State struct {
	count int32
}

func (s *State) setState(i int) {
	atomic.AddInt32(&s.count, int32(i))
}

func main() {
}
