package main

import "fmt"

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
	s.loop() // this will block
}

func (s *Server) loop() {
	for {
		// selects are specifically for channels
		select {
		case <-s.quitch:
			// do some stuff when we need to quit
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
			// this is just to satisfy the compiler... obviously, it does nothing
		}
	}
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("message received:", msg)
}

func main() {
	server := newServer()
	go server.start() // schedule as a goroutine

	// one thing we could do is pipe in a message
	server.msgch <- "hey do this!"
}
