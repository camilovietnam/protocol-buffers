package main

import (
	"context"
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
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("[i] call Greet(): %v", req)
	g := req.GetGreeting()
	return &greetpb.GreetResponse{
		Result: fmt.Sprintf("Buenos dias, %s %s", g.GetFirstName(), g.GetLastName()),
	}, nil
}
