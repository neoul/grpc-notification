package main

import (
	"log"
	"net"
	"os"

	noti "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
)

const (
	port       = ":50051"
	serverName = "notification-server"
)

type notificationServer struct {
	serverName string
	client     map[string]noti.Notification_SubscribeServer
}

// Register RPC implementation

func (notiServer *notificationServer) Subscribe(req *noti.Subscription, srv noti.Notification_SubscribeServer) error {
	log.Printf("Received Subscription: %v", req)
	notiServer.client[req.GetName()] = srv
	if err := srv.Send(&noti.Notification{Message: "HI " + notiServer.serverName}); err != nil {
		// log.Fatalf("Send failed %v", err)
		log.Printf("Send failed %v", err)
		return err
	}
	return nil
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
	noti.RegisterNotificationServer(grpcServer, &notificationServer{serverName: name, client: make(map[string]noti.Notification_SubscribeServer)})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
