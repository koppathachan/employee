package main

import (
	"log"
	"net"

	"github.com/sasidakh/employee/people"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	people.RegisterPeopleServer(s, people.Server{})
	log.Println(listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
