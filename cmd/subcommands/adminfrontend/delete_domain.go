package adminfrontend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webimizer.dev/poem/admin"
)

/* gRPC client DeleteDomain */
func (p *adminFrontendCmd) grpcDeleteDomain(req *admin.AdminDomain) (response *admin.DomainResponse, err error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.gRPCAdminHost, p.gRPCAdminPort), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := admin.NewAdminClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.DeleteDomain(ctx, req)
}
