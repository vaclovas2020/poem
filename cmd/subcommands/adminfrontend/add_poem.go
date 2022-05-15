package adminfrontend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webimizer.dev/poem/admin"
)

func (p *adminFrontendCmd) grpcAddPoem(req *admin.AdminPoem) (response *admin.PoemResponse, err error) {
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
	return client.AddPoem(ctx, req)
}
