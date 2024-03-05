package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/arimura/grpc_intro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ToDoServer struct {
	grpc_intro.UnimplementedToDoServer
}

func (s *ToDoServer) GetItem(context.Context, *grpc_intro.Cond) (*grpc_intro.Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
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
