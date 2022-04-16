package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"voting-server/votingpb"
)

func main() {
	dialer, err := grpc.Dial("0.0.0.0:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[x] Dial error: %v", err)
	}

	defer func() {
		_ = dialer.Close()
	}()

	votingClient := votingpb.NewVotingServiceClient(dialer)
	stream := getVotingStream(votingClient)

	vote(stream, 1)
	vote(stream, 2)
	vote(stream, 2)
	vote(stream, 3)
	vote(stream, 2)
	response := closeVoting(stream)

	// Winner should be option 2!
	fmt.Println("[>] Winner was: Option ", response.GetWinner())
}

func getVotingStream(client votingpb.VotingServiceClient) votingpb.VotingService_VoteClient {
	stream, err := client.Vote(context.Background())
	if err != nil {
		log.Fatalf("[x] Vote error: %v", err)
	}
	return stream
}

func vote(stream votingpb.VotingService_VoteClient, option uint32) {
	if err := stream.Send(&votingpb.VoteRequest{Option: option}); err != nil {
		log.Fatalf("[x] Send error: %v", err)
	}
}

func closeVoting(stream votingpb.VotingService_VoteClient) *votingpb.VoteResponse {
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[x] CloseAndRecv error: %v", err)
	}
	return response
}
