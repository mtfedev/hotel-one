package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 0 bytes
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server sterting")
	s.loop() // block
}

func (s *Server) quit() {
	s.quitch <- struct{}{}
}
func (s *Server) sengMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch: // do some dtuff when we need to quit
			fmt.Println("quiting server")
			break mainloop
		case msg := <-s.msgch:
			s.handelMessage(msg) // do some dtuff when we have a message
		default:
		}
	}
	fmt.Println("server is shutting donw gracefully")
}
func (s *Server) handelMessage(msg string) {
	fmt.Println("we received a message", msg)
}
func main() {
	server := newServer()
	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()

}
