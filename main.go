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
	// result := fetchResource()
	// fmt.Println("result:", result)

	// create go routine (non-blocking call)
	// go fetchResource()
	// at this point in the code execution, fetchResource() has not completed

	// // schedule anonymous function -> identical to "go fetchResource()"
	// go func() {
	// 	fetchResource()
	// }()

	// this won't work -> where is the println? it's nowhere.. because it's running async -> threads getting threads
	go func() {
		result := fetchResource(1)
		fmt.Println(result)
	}()
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
