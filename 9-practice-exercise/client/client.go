package main

import (
	"context"
	"io"
	"log"
	"os"
	"prime-decomposition-server/primepb"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		err error
		ctx = context.Background()
	)

	con, err := grpc.Dial("0.0.0.0:51234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[x] dial: %v", err)
	}

	defer func() {
		_ = con.Close()
	}()

	number, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("[x] parseInt: %v", err)
	}

	client := primepb.NewPrimeServiceClient(con)
	stream, err := client.Decompose(ctx, &primepb.DecomposeRequest{
		Number: number,
	})
	if err != nil {
		log.Fatalf("[x], decompose: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			log.Print("[!] end of stream")
			return
		}

		if err != nil {
			log.Fatalf("[x] recv: %v", err)
		}

		factor := res.GetFactor()
		log.Printf("[>] Returned factor: %v", factor)
	}
}
