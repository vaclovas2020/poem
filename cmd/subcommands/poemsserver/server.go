package poemsserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"webimizer.dev/poem/poems"
)

type poemsServer struct {
	poems.UnimplementedPoemsServer
}

func runServer(host string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	poems.RegisterPoemsServer(grpcServer, &poemsServer{})
	log.Printf("Starting listen on %s:%d", host, port)
	log.Fatal(grpcServer.Serve(lis))
}
