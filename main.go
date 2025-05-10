package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"neko-acm-client/pkg/pb"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer ",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	dial, err := grpc.Dial("127.0.0.1:14516", opts...)
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	problemClient := pb.NewProblemServiceClient(dial)

	pi := &pb.ProblemInstructionRequest{
		Title: "Hello World",
	}

	resp, err := problemClient.GenerateProblem(context.Background(), pi)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", resp)
}
