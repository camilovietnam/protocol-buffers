package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer func() {
		_ = con.Close()
	}()

	cli := greetpb.NewGreetServiceClient(con)
	// greet(cli)
	greetManyTimes(cli)
}

func greet(cli greetpb.GreetServiceClient) {
	res, err := cli.Greet(context.Background(), &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dani",
			LastName:  "Filth",
		},
	})

	if err != nil {
		log.Fatalf("greet: %#v", err)
	}

	fmt.Printf("[>] %s\n", res.GetResult())
}

func greetManyTimes(cli greetpb.GreetServiceClient) {
	var (
		stream greetpb.GreetService_GreetManyTimesClient
		req    = &greetpb.GreetManyTimesRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Hector",
				LastName:  "Lavoe",
			},
		}
	)
	stream, err := cli.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("[x] greetManyTimes: %#v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("[<] close stream")
			return
		}
		if err != nil {
			log.Fatalf("[x] recv: %#v", err)
		}

		fmt.Printf("[>] %s\n", res.GetResult())
	}

}
