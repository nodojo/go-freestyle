package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1) // does some work
		go func(i int) {
			state.count = i + 1
			wg.Done() // done working
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", state)
}
