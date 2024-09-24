package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Deepanshuisjod/rest-api/protos/auth" // Ensure path matches
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedAuthServer
}

// Correct method name to match the gRPC service definition
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetMessage()) // Correct field name if 'Name' is used in .proto
	return &pb.HelloResponse{Message: "Hello " + in.GetMessage()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
