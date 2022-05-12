/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"webimizer.dev/poem/admin"
)

/* Admin gRPC server */
type adminServer struct {
	admin.UnimplementedAdminServer
	cmd *adminServerCmd // admin server subcommand struct
}

/* Run server listener */
func (p *adminServerCmd) runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.host, p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	admin.RegisterAdminServer(grpcServer, &adminServer{cmd: p})
	log.Printf("Starting listen on %s:%d", p.host, p.port)
	log.Fatal(grpcServer.Serve(lis))
}
