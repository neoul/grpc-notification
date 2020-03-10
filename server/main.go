package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	noti "github.com/neoul/grpc-notification/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// stdin -> select -> send
// recv  ->

const (
	port = ":50051"
	name = "notification-server"
)

type notitificationClient struct {
	name   string
	stream noti.Notification_SubscribeServer
}

type notificationServer struct {
	name       string
	enabled    chan bool
	grpcServer *grpc.Server
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
			if err == io.EOF {
				return nil
			}
			log.Printf("%v\n", err)
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
	// var done bool = false
	for {
		fmt.Printf("%s> ", notiServer.name)
		fmt.Scan(&in)
		// log.Println(in, strings.HasPrefix(in, "qu"))
		if strings.HasPrefix(in, "qu") {
			notiServer.grpcServer.GracefulStop()
			notiServer.enabled <- true
			return
		}
		for k, client := range notiServer.client {
			log.Printf("Send notification to %s", k)
			if err := client.stream.Send(&noti.Notification{Message: in}); err != nil {
				log.Fatalf("Send failed %v", err)
			}
		}
	}
}

func main() {
	encrypt := flag.Bool("encrypt", false, "enable encryption of gRPC")
	certfile := flag.String("certfile", "", "'server.pem (server.crt)' server certificate (public key)")
	keyfile := flag.String("keyfile", "", "'server.key' server private key")
	flag.Usage = func() {
		fmt.Printf(" %s <FLAG> <SERVER_NAME>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	name := name
	if len(args) > 0 {
		name = args[0]
	}
	fmt.Printf("Server starts with '%s'\n", name)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var grpcServer *grpc.Server
	if *encrypt {
		creds, err := credentials.NewServerTLSFromFile(*certfile, *keyfile)
		if err != nil {
			log.Fatalf("failed to load TLS: %v", err)
		}
		grpcServer = grpc.NewServer(grpc.Creds(creds))
		fmt.Println("Server start within encryption mode ..")
	} else {
		grpcServer = grpc.NewServer()
	}
	notiServer := &notificationServer{
		name:       name,
		enabled:    make(chan bool),
		grpcServer: grpcServer,
		client:     make(map[string]*notitificationClient)}
	noti.RegisterNotificationServer(grpcServer, notiServer)
	go notiServer.Notify()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	<-notiServer.enabled
}
