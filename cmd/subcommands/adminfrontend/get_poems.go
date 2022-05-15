package adminfrontend

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webimizer.dev/poem/poems"
)

func (p *adminFrontendCmd) grpcGetPoems(req *poems.PoemsRequest) (response *poems.PoemsResponse, err error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.gRPCPoemsHost, p.gRPCPoemsPort), opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := poems.NewPoemsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.GetPoems(ctx, req)
}
