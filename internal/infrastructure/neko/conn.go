package neko

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"nekoacm-client/pkg/config"
)

var (
	ClientConn *grpc.ClientConn
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + config.Conf.NekoGrpc.Token,
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

// 创建 gRPC 连接
func createConn(address string) error {
	var err error
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	// 创建gRPC连接
	ClientConn, err = grpc.Dial(address, opts...)
	if err != nil {
		return err
	}

	return nil
}
