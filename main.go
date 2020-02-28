package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wolfmib/user_grpc_v1/user_proto"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

// server is used to implement helloworld.
type server struct {
	pb.UnimplementedUserServiceServer
}

// RegisterApi implements UserServicesServer
func (s *server) RegisterApi(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received: %v  ", in.GetFirstName())
	log.Printf("Received: %v  ", in.GetFamilyName())
	log.Printf("Received: %v  ", in.GetEmail())
	return &pb.RegisterResponse{Uuid: "cest uuid", Email: "xxx@gmail.com", UserId: 31}, nil
}

func main() {

	log.Println("[Jean]: I am backend.")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
