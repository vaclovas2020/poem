package oauthserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"webimizer.dev/poem/oauth"
)

/* gRPC AuthUser */
func (srv *oAuthServer) AuthUser(_ context.Context, request *oauth.AuthRequest) (response *oauth.AuthResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}
