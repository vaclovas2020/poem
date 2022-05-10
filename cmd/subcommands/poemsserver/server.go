/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package poemsserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"webimizer.dev/poem/poems"
)

/* Poems gRPC server */
type poemsServer struct {
	poems.UnimplementedPoemsServer
	cmd *poemsServerCmd // poems server subcommand struct
}

/* Run server listener */
func (p *poemsServerCmd) runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.host, p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	poems.RegisterPoemsServer(grpcServer, &poemsServer{cmd: p})
	log.Printf("Starting listen on %s:%d", p.host, p.port)
	log.Fatal(grpcServer.Serve(lis))
}
