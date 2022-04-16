package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"voting-server/votingpb"
)

type server struct{}
type candidate = uint32

func calculateWinner(votes map[candidate]int) uint32 {
	maxVotes := 0
	var winner candidate = 0

	for option, optionVotes := range votes {
		if optionVotes > maxVotes {
			winner = option
			maxVotes = optionVotes
		}
	}

	return winner
}

func notifyClient(stream votingpb.VotingService_VoteServer, winner uint32) {
	err := stream.SendAndClose(&votingpb.VoteResponse{
		Winner: winner,
	})

	if err != nil {
		log.Fatalf("[x] SendMsg error: %v", err)
	}
}
func (s *server) Vote(stream votingpb.VotingService_VoteServer) error {
	var votes = make(map[candidate]int)

	for {
		req, err := stream.Recv()

		// if I get the EOF signal, I count the votes and inform the result
		if err == io.EOF {
			fmt.Println("[>] Time to count the votes")

			winner := calculateWinner(votes)
			notifyClient(stream, winner)

			fmt.Println("[>] The winner is: Option ", winner)
			break
		}

		// if I get a different error, I exit with Fatal
		if err != nil {
			log.Fatalf("[x] Recv error: %v", err)
		}

		// if I get a vote, I count it
		option := req.GetOption()
		votes[option]++
		fmt.Println("[+] One vote for: ", option)
	}

	return nil
}

func main() {
	// create the net listener
	listener, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatalf("[x] listen: %v", err)
	}

	// create the grpc server and associate the server structure with the implemented methods
	s := grpc.NewServer()
	votingpb.RegisterVotingServiceServer(s, &server{})

	// run the grpc server on the listener
	log.Printf("[!] Server running in port 12345")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("serve: %v", err)
	}
}
