package main

import (
	"bidi/bidipb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	// create a dialer
	dialer, err := grpc.Dial("0.0.0.0:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[x] Dial error: %v", err)
	}
	defer func() {
		_ = dialer.Close()
	}()

	// associate a new client to dialer
	client := bidipb.NewGreetServiceClient(dialer)
	stream := getStream(client)

	// Send messages to the server using a go routine
	go func() {
		time.Sleep(2222 * time.Millisecond)
		sendUsername(stream, "Alice")
		time.Sleep(2222 * time.Millisecond)
		sendUsername(stream, "Bob")
		time.Sleep(2222 * time.Millisecond)
		sendUsername(stream, "Charlie")
		err = stream.CloseSend()
	}()

	// now sit back and wait for incoming messages from the server
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("[>] Server shut down the connection")
			break
		}

		if err != nil {
			log.Fatalf("Recv error: %v", err)
		}

		result := res.GetResult()
		fmt.Println("[>] Server said: ", result)
	}
}

func sendUsername(stream bidipb.GreetService_GreetEveryoneClient, username string) {
	err := stream.Send(&bidipb.GreetEveryoneRequest{Username: username})
	if err != nil {
		log.Fatalf("[x] Send error: %v", err)
	}
}
func getStream(client bidipb.GreetServiceClient) bidipb.GreetService_GreetEveryoneClient {
	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("[x] GreetEveryone error: %v", err)
	}
	return stream
}
