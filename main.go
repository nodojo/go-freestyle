package main

import (
	"fmt"
	"time"
)

// the basic reason golang was designed:
// concurrency, distributed systems, asynchronous environments
func main() {
	// go routine is a function that is being scheduled by the go scheduler (in the go runtime... each time you build a binary, each time you build your own program, it's going to compile in the golang runtime)...
	// the scheduler takes all your goroutines and it will make sure that if a goroutine has no task to perform, it's going to yield the processor and it's going to schedule something else...
	// understand that in go, there is no such thing as parallelism (goroutines don't run in parallel), it's basically concurrency
	result := fethcResource()
	fmt.Println("result:", result)
}

func fethcResource() string {
	time.Sleep(time.Second * 2)
	return "some result"
}
