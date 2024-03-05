package main

import (
	"context"
	"log"
	"time"

	"github.com/arimura/grpc_intro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := grpc_intro.NewToDoClient(conn)
	d := int32(1)
	cond := &grpc_intro.Cond{Day: &d}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := client.GetItem(ctx, cond)
	if err != nil {
		log.Fatalf("fail to call GetItem: %v", err)
	}
	log.Printf("Response: %v", r)
}
