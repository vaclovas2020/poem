package adminfrontend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webimizer.dev/poem/oauth"
)

func (p *adminFrontendCmd) grpcAuthUser(user *oauth.AuthRequest) (response *oauth.AuthResponse, err error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.gRPCOauthHost, p.gRPCOauthPort), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := oauth.NewOauthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.AuthUser(ctx, user)
}
