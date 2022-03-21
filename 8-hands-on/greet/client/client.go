package main

import (
	"context"
	"fmt"
	"greet/greetpb"
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
	greet(cli)
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
