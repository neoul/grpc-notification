package main

import (
	"context"
	"log"
	"os"

	noti "github.com/neoul/grpc-notification/proto"
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
	notiClient := noti.NewNotificationClient(conn)

	// Contact the server and print out its response.
	ctx := context.Background()
	subClient, err := notiClient.Subscribe(ctx)
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}
	if err := subClient.Send(&noti.Subscription{Name: name}); err != nil {
		log.Fatalf("could not send: %v", err)
	}
	for {
		noti, err := subClient.Recv()
		if err != nil {
			log.Fatalf("could not recv: %v", err)
		}
		log.Println(noti.GetMessage())
	}
}
