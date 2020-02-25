package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	noti "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
)

// stdin -> select -> send
// recv  ->

const (
	port       = ":50051"
	serverName = "notification-server"
)

type notitificationClient struct {
	name   string
	stream noti.Notification_SubscribeServer
}

type notificationServer struct {
	serverName string
	client     map[string]*notitificationClient
}

func (notiServer *notificationServer) Subscribe(srv noti.Notification_SubscribeServer) error {
	for {
		in, err := srv.Recv()
		if err != nil {
			for key, client := range notiServer.client {
				if client.stream == srv {
					delete(notiServer.client, key)
					break
				}
			}
			log.Printf("%v\n", err)
			if err == io.EOF {
				return nil
			}
			return err
		}
		name := in.GetName()
		log.Println("Connected", name)
		if notiServer.client[name] != nil && notiServer.client[name].stream != srv {
			notiServer.client[name].stream.Context().Done()
		}
		notiServer.client[name] = &notitificationClient{name: name, stream: srv}
	}
}

func (notiServer *notificationServer) Notify() {
	var in string
	for {
		fmt.Printf("IN> ")
		fmt.Scan(&in)
		// time.Sleep(1 * time.Second)
		for k, client := range notiServer.client {
			log.Printf("Send notification to %s", k)
			if err := client.stream.Send(&noti.Notification{Message: in}); err != nil {
				log.Fatalf("Send failed %v", err)
			}
		}
	}
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
	grpcServer := grpc.NewServer()
	notiServer := &notificationServer{serverName: name, client: make(map[string]*notitificationClient)}
	noti.RegisterNotificationServer(grpcServer, notiServer)
	go notiServer.Notify()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
