package main

import (
	"fmt"
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
	// go func() {
	// 	result := fetchResource(1)
	// 	fmt.Println(result)
	// }()

	// think of channels as a tunnel, and people on the other side of the tunnel will receive what you put into the tunnel
	// two types of channels (it's a common practice to append "ch" to any variable that represents a channel)
	// 1. unbuffered channel
	// 2. buffered channel
	// **note: a channel in golang will always block if it's full
	// THIS EXAMPLE CREATES A DEADLOACK BECAUSE IT'S UNBUFFERED WHICH MEANS IT IS FULL WHEN WE SEND "foo" TO IT
	resultch := make(chan string) // -> make me a channel of type string (unbuffered) // -> holds ONLY whatever is sent to it, in this case "foo" *which means it will block, because it's full
	// resultch := make(chan string, 10) // -> make me a channel of type string (buffered with a buffer size of 10) -> holds 10
	// resultch <- "foo"
	// result := <-resultch
	// fmt.Println(result)
	// THIS CODE WILL NEVER EXECUTE BECAUSE OF THE DEADLOCK CREATED ABOVE!

	// so either buffer the channel, or create a goroutine to read from the channel (below)
	go func() {
		result := <-resultch
		fmt.Println(result)
	}()

	resultch <- "foo"
}

// func fetchResource(n int) string {
// 	time.Sleep(time.Second * 2)
// 	return fmt.Sprintf("result %d", n)
// }
