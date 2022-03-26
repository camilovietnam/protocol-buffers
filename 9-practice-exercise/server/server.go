package main

import (
	"log"
	"net"
	"prime-decomposition-server/primepb"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	var err error

	listener, err := net.Listen("tcp", "0.0.0.0:51234")
	if err != nil {
		log.Fatalf("[x] listen: %v", err)
	}

	s := grpc.NewServer()

	log.Printf("[!] Server running in port 51234")

	primepb.RegisterPrimeServiceServer(s, &server{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("[x] serve: %v", err)
	}

}

func (s *server) Decompose(req *primepb.DecomposeRequest, stream primepb.PrimeService_DecomposeServer) error {
	log.Printf("[!] decompose request: %d", req.GetNumber())
	num := req.GetNumber()
	var factor int64 = 2

	for num > 1 {
		if num%factor == 0 {
			if err := stream.Send(&primepb.DecomposeResponse{
				Factor: factor,
			}); err != nil {
				return err
			}
			num /= factor
		} else {
			factor++
		}
	}

	log.Printf("[^] end of request")

	return nil
}
