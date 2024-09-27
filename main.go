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
		msgch:  make(chan string),
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
			// do some stuff when we have a message
			_ = msg // this prevents the compiler from complaining
		}
	}
}

func main() {

}
