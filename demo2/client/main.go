package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/arimura/grpc-intro/demo2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up connection to server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	client := pb.NewHogeServiceClient(conn)

	// Contact the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get ID from command line or use default
	id := "1"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	// Call GetHoge
	res, err := client.GetHoge(ctx, &pb.GetHogeRequest{Id: id})
	if err != nil {
		log.Fatalf("GetHoge failed: %v", err)
	}

	// Print the response
	hoge := res.GetHoge()
	log.Printf("Hoge: ID=%s, Name=%s, Description=%s",
		hoge.GetId(), hoge.GetName(), hoge.GetDescription())
}
