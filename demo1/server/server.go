package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/arimura/grpc_intro"
	"google.golang.org/grpc"
)

type ToDoServer struct {
	grpc_intro.UnimplementedToDoServer
}

func (s *ToDoServer) GetItem(context.Context, *grpc_intro.Cond) (*grpc_intro.Item, error) {
	c := "Hello, World!"
	r := &grpc_intro.Item{
		Content: &c,
	}
	return r, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpc_intro.RegisterToDoServer(grpcServer, &ToDoServer{})
	grpcServer.Serve(lis)
}
