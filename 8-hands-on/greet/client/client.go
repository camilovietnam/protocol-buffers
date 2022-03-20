package main

import (
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
	defer con.Close()

	cli := greetpb.NewGreetServiceClient(con)

	fmt.Printf("Created client %#v", cli)
}
