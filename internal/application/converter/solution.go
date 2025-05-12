package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/pkg/pb"
)

// SolutionInstructionToRequest 将题解指令DTO转换为gRPC请求
func SolutionInstructionToRequest(instruction dto.SolutionInstruction) *pb.SolutionInstructionRequest {
	return &pb.SolutionInstructionRequest{
		Title:        instruction.Title,
		Description:  instruction.Description,
		Input:        instruction.Input,
		Output:       instruction.Output,
		SampleInput:  instruction.SampleInput,
		SampleOutput: instruction.SampleOutput,
		Hint:         instruction.Hint,
		Language:     instruction.Language,
	}
}

// RequestToSolutionInstruction 将gRPC请求转换为题解指令DTO
func RequestToSolutionInstruction(request *pb.SolutionInstructionRequest) dto.SolutionInstruction {
	return dto.SolutionInstruction{
		Title:        request.Title,
		Description:  request.Description,
		Input:        request.Input,
		Output:       request.Output,
		SampleInput:  request.SampleInput,
		SampleOutput: request.SampleOutput,
		Hint:         request.Hint,
		Language:     request.Language,
	}
}

// SolutionToResponse 将题解DTO转换为gRPC响应
func SolutionToResponse(solution dto.Solution) *pb.SolutionResponse {
	return &pb.SolutionResponse{
		SourceCode:  solution.SourceCode,
		Explanation: solution.Explanation,
		Language:    solution.Language,
	}
}

// ResponseToSolution 将gRPC响应转换为题解DTO
func ResponseToSolution(resp *pb.SolutionResponse) dto.Solution {
	return dto.Solution{
		SourceCode:  resp.SourceCode,
		Explanation: resp.Explanation,
		Language:    resp.Language,
	}
}
