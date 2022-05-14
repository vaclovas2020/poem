/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package oauthserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"webimizer.dev/poem/oauth"
)

/* Admin gRPC server */
type oAuthServer struct {
	oauth.UnimplementedOauthServer
	cmd *oauthServerCmd // oauth server subcommand struct
}

/* Run server listener */
func (p *oauthServerCmd) runServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", p.host, p.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	oauth.RegisterOauthServer(grpcServer, &oAuthServer{cmd: p})
	log.Printf("Starting listen on %s:%d", p.host, p.port)
	log.Fatal(grpcServer.Serve(lis))
}
