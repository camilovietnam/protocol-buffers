package main

import (
	"bidi/bidipb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
}

func (s *server) GreetEveryone(stream bidipb.GreetService_GreetEveryoneServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("[>] Server received end of stream.")
			return nil
		}

		if err != nil {
			log.Fatalf("[x] Recv error: %v", err)
		}

		username := request.GetUsername()
		fmt.Println("[>] Server got username: ", username)
		err = stream.Send(&bidipb.GreetEveryoneResponse{
			Result: "[i] Wassup " + username + ", welcome!",
		})

		if err != nil {
			log.Fatalf("[x] Send error: %v", err)
		}
	}

	return nil
}

func main() {
	// create the listener
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatalf("[x] Listen error: %v", err)
	}

	// create the grpc server
	s := grpc.NewServer()

	// attach the structure to the server
	bidipb.RegisterGreetServiceServer(s, &server{})

	// listen the grpc on the listener
	fmt.Println("[>] Server is running in port 12345.")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("[x] Serve error: %v", err)
	}
}
