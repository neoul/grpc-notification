package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:50051"
	clientName = "notification-client"
)

func main() {
	name := clientName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNotificationClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Register(ctx, &pb.RegistrationRequest{ClientName: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting %s by %s", r.GetClientName(), r.GetServerName())

	c.Subscribe(ctx, &pb.SubscribeRequest{ClientName: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting %s by %s", r.GetClientName(), r.GetServerName())
}
