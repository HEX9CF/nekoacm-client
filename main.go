package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"neko-acm-client/pkg/pb"
)

func main() {
	dial, err := grpc.Dial("127.0.0.1:14516", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
