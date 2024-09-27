package main

import (
	"fmt"
	"time"
)

type Server struct {
	// quitch chan bool  // will allocate 1 byte
	// trick used for channels that will just need to communicate some binary expression (ex: true or false)
	quitch chan struct{} // will allocate 0 bytes
	msgch  chan string
}

// this is a constructor -> good practice to signify by prepending "new" to the name
func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128), // now we buffer the msg channel
	}
}

func (s *Server) start() {
	fmt.Println("server starting")
	s.loop()
}

// func (s *Server) sendMessage(msg string) {
// 	s.msgch <- msg
// }

func (s *Server) quit() {
	// could do this...
	// close(s.quitch)
	// ...or this
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
	// go allows you to declare a loop (here our declared loop is named "mainloop")
mainloop:
	for {
		// selects are specifically for channels
		select {
		case <-s.quitch: // when quit() is called, it will cause this case to get triggered
			fmt.Println("quitting server")
			break mainloop // this allows us to trigger a break on the outer for loop when this case is satisfied
		case msg := <-s.msgch:
			s.handleMessage(msg)
		}
	}
	fmt.Println("server is shutting down gracefully")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("message received:", msg)
}

func main() {
	server := newServer()

	// add func to spin up another goroutine
	// this allows code execution to continue without blocking our `go server.start()` line below
	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()

	// // one thing we could do is pipe in a message
	// // server.msgch <- "hey do this!"
	// for i := 0; i < 100; i++ {
	// 	server.sendMessage(fmt.Sprintf("handle this number %d", i))
	// }
	// // another way would be to call a function that performs the same action
	// server.sendMessage("hey! do this...")
	// // since everything is happening asynchronously, pause the program so that we can see it working
	// time.Sleep(time.Second * 5)
}
