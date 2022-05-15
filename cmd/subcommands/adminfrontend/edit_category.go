package adminfrontend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webimizer.dev/poem/admin"
)

/* gRPC client EditCategory */
func (p *adminFrontendCmd) grpcEditCategory(req *admin.AdminCategoryEdit) (response *admin.CategoryEditResponse, err error) {
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
	return client.EditCategory(ctx, req)
}
