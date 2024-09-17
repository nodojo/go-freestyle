package main

import (
	"fmt"
)

func main() {
	// if you write to a channel faster than your reading from it, at some point it will become full
	// in programming, when you are specifying a size allocation, use values that are the power of 2 (these are values a computer can understand easily)
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch)

	// // this only reads the first input ("A") from the channel
	// msg := <-msgch
	// fmt.Println("msg is:", msg)
	// // this will read the next input ("B")
	// msg = <-msgch
	// fmt.Println("msg is:", msg)

	// range over a channel (this is our consumer)
	// this will display all inputs, but throw a "fatal error: all goroutines are asleep - deadlock" error...
	// .. this is because no stopping opint has been set
	// for msg := range msgch {
	// 	fmt.Println(msg)
	// }

	// create never-ending for loop (behaves like a while loop)
	// BAD! -> this will continue iterating forever (this case is automatically handled by using a range)
	for {
		msg := <-msgch
		fmt.Println("message ->", msg)
	}

	fmt.Println("done reading all the messages from the channel!")
}

// func fetchResource(n int) string {
// 	time.Sleep(time.Second * 2)
// 	return fmt.Sprintf("result %d", n)
// }
