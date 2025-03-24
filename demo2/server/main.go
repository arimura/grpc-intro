package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/arimura/grpc-intro/demo2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedHogeServiceServer
	hogeDB map[string]*pb.Hoge
}

func newServer() *server {
	s := &server{
		hogeDB: make(map[string]*pb.Hoge),
	}
	// Pre-populate with some data
	s.hogeDB["1"] = &pb.Hoge{
		Id:          "1",
		Name:        "First Hoge",
		Description: "This is the first hoge item",
	}
	s.hogeDB["2"] = &pb.Hoge{
		Id:          "2",
		Name:        "Second Hoge",
		Description: "This is the second hoge item",
	}
	return s
}

func (s *server) GetHoge(ctx context.Context, req *pb.GetHogeRequest) (*pb.GetHogeResponse, error) {
	log.Printf("GetHoge request for ID: %s", req.Id)

	hoge, exists := s.hogeDB[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Hoge with ID '%s' not found", req.Id)
	}

	return &pb.GetHogeResponse{
		Hoge: hoge,
	}, nil
}

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHogeServiceServer(grpcServer, newServer())

	log.Printf("Server listening on port %d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
