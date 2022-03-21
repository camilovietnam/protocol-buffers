package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"log"
	"net"
	"time"

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

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("[i] call GreetManyTimes(): %v", req)
	var (
		g         = req.GetGreeting()
		greetings = [...]string{"Hola", "Hallo", "Xin Chao", "Konnichiwa"}
	)

	for _, s := range greetings {
		stream.Send(&greetpb.GreetManyTimesResponse{
			Result: fmt.Sprintf("%s, %s %s!", s, g.GetFirstName(), g.GetLastName()),
		})
		time.Sleep(1200 * time.Millisecond)
	}
	return nil
}
