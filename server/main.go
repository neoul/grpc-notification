package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
)

const (
	port       = ":50051"
	serverName = "notification-server"
)

type notificationServer struct {
	name string
	pb.UnimplementedNotificationServer
}

// Register RPC implementation
func (s *notificationServer) Register(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationResponse, error) {
	return &pb.RegistrationResponse{ClientName: req.GetClientName(), ServerName: s.name}, nil
}

func main() {
	name := serverName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificationServer(s, &notificationServer{name: name})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
