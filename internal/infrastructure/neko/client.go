package neko

import (
	"nekoacm-client/pkg/pb"
)

// 全局客户端变量
var (
	AssistantClient pb.AssistantServiceClient
	JudgeClient     pb.JudgeServiceClient
	ProblemClient   pb.ProblemServiceClient
	SolutionClient  pb.SolutionServiceClient
	TestcaseClient  pb.TestcaseServiceClient
	MiscClient      pb.MiscServiceClient
)

// 初始化各个服务的客户端
func initClient() {
	AssistantClient = pb.NewAssistantServiceClient(ClientConn)
	JudgeClient = pb.NewJudgeServiceClient(ClientConn)
	ProblemClient = pb.NewProblemServiceClient(ClientConn)
	SolutionClient = pb.NewSolutionServiceClient(ClientConn)
	TestcaseClient = pb.NewTestcaseServiceClient(ClientConn)
	MiscClient = pb.NewMiscServiceClient(ClientConn)
}
