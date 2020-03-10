package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	noti "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
	name    = "notification-client"
)

func main() {
	encrypt := flag.Bool("encrypt", false, "enable encryption of gRPC")
	certfile := flag.String("certfile", "", "'ca.pem (ca.crt) or server.pem (server.crt)' (server certificate or CA certificate)")
	flag.Usage = func() {
		fmt.Printf(" %s <FLAG> <CLIENT_NAME>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	name := name
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Client starts with '%s'\n", name)

	var opt grpc.DialOption
	if *encrypt {
		creds, err := credentials.NewClientTLSFromFile(*certfile, "")
		if err != nil {
			log.Fatalf("failed to load TLS: %v", err)
		}
		opt = grpc.WithTransportCredentials(creds)
	} else {
		opt = grpc.WithInsecure()
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opt, grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	notiClient := noti.NewNotificationClient(conn)

	// Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()
	ctx := context.Background()
	subClient, err := notiClient.Subscribe(ctx)
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}
	if err := subClient.Send(&noti.Subscription{Name: name}); err != nil {
		log.Fatalf("failed to send: %v", err)
	}
	for {
		noti, err := subClient.Recv()
		if err != nil {
			log.Fatalf("failed to recv: %v", err)
		}
		log.Println(noti.GetMessage())
	}
}
