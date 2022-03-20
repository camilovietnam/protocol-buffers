package main

import (
	"fmt"
	"greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("serve: %v", err)
	}

	fmt.Println("Server is listening in port: 50051")
}
