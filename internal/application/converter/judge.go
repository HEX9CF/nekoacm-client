package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/pkg/pb"
)

// SubmissionToRequest 将提交信息DTO转换为gRPC请求
func SubmissionToRequest(submission dto.Submission) *pb.SubmissionRequest {
	return &pb.SubmissionRequest{
		SourceCode:     submission.SourceCode,
		Language:       submission.Language,
		Stdin:          submission.Stdin,
		ExpectedOutput: submission.ExpectedOutput,
	}
}

// RequestToSubmission 将gRPC请求转换为提交信息DTO
func RequestToSubmission(request *pb.SubmissionRequest) dto.Submission {
	return dto.Submission{
		SourceCode:     request.SourceCode,
		Language:       request.Language,
		Stdin:          request.Stdin,
		ExpectedOutput: request.ExpectedOutput,
	}
}

// JudgementToResponse 将评测结果DTO转换为gRPC响应
func JudgementToResponse(judgement dto.Judgement) *pb.JudgementResponse {
	return &pb.JudgementResponse{
		Stdout:        judgement.Stdout,
		Stderr:        judgement.Stderr,
		CompileOutput: judgement.CompileOutput,
		Message:       judgement.Message,
		Status:        judgement.Status,
	}
}

// ResponseToJudgement 将gRPC响应转换为评测结果DTO
func ResponseToJudgement(resp *pb.JudgementResponse) dto.Judgement {
	return dto.Judgement{
		Stdout:        resp.Stdout,
		Stderr:        resp.Stderr,
		CompileOutput: resp.CompileOutput,
		Message:       resp.Message,
		Status:        resp.Status,
	}
}
