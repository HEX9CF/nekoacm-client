package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-common/api/proto/pb"
)

// ProblemInstructionToRequest 将题目指令DTO转换为gRPC请求
func ProblemInstructionToRequest(instruction dto.ProblemInstruction) *pb.ProblemInstructionRequest {
	return &pb.ProblemInstructionRequest{
		Title:        instruction.Title,
		Description:  instruction.Description,
		Input:        instruction.Input,
		Output:       instruction.Output,
		SampleInput:  instruction.SampleInput,
		SampleOutput: instruction.SampleOutput,
		Hint:         instruction.Hint,
		Tags:         instruction.Tags,
	}
}

// RequestToProblemInstruction 将gRPC请求转换为题目指令DTO
func RequestToProblemInstruction(request *pb.ProblemInstructionRequest) dto.ProblemInstruction {
	return dto.ProblemInstruction{
		Title:        request.Title,
		Description:  request.Description,
		Input:        request.Input,
		Output:       request.Output,
		SampleInput:  request.SampleInput,
		SampleOutput: request.SampleOutput,
		Hint:         request.Hint,
		Tags:         request.Tags,
	}
}

// ProblemToResponse 将题目DTO转换为gRPC响应
func ProblemToResponse(problem dto.Problem) *pb.ProblemResponse {
	return &pb.ProblemResponse{
		Title:        problem.Title,
		Description:  problem.Description,
		Input:        problem.Input,
		Output:       problem.Output,
		SampleInput:  problem.SampleInput,
		SampleOutput: problem.SampleOutput,
		Hint:         problem.Hint,
		Tags:         problem.Tags,
	}
}

// ResponseToProblem 将gRPC响应转换为题目DTO
func ResponseToProblem(resp *pb.ProblemResponse) dto.Problem {
	return dto.Problem{
		Title:        resp.Title,
		Description:  resp.Description,
		Input:        resp.Input,
		Output:       resp.Output,
		SampleInput:  resp.SampleInput,
		SampleOutput: resp.SampleOutput,
		Hint:         resp.Hint,
		Tags:         resp.Tags,
	}
}

// ProblemDataToRequest 将题目数据DTO转换为gRPC请求
func ProblemDataToRequest(data dto.ProblemData) *pb.ProblemDataRequest {
	return &pb.ProblemDataRequest{
		Data: data.Data,
	}
}

// RequestToProblemData 将gRPC请求转换为题目数据DTO
func RequestToProblemData(request *pb.ProblemDataRequest) dto.ProblemData {
	return dto.ProblemData{
		Data: request.Data,
	}
}

// TranslateInstructionToRequest 将翻译指令DTO转换为gRPC请求
func TranslateInstructionToRequest(instruction dto.TranslateInstruction) *pb.TranslateInstructionRequest {
	return &pb.TranslateInstructionRequest{
		Title:       instruction.Title,
		Description: instruction.Description,
		Input:       instruction.Input,
		Output:      instruction.Output,
		Hint:        instruction.Hint,
		Tags:        instruction.Tags,
		TargetLang:  instruction.TargetLang,
	}
}

// RequestToTranslateInstruction 将gRPC请求转换为翻译指令DTO
func RequestToTranslateInstruction(request *pb.TranslateInstructionRequest) dto.TranslateInstruction {
	return dto.TranslateInstruction{
		Title:       request.Title,
		Description: request.Description,
		Input:       request.Input,
		Output:      request.Output,
		Hint:        request.Hint,
		Tags:        request.Tags,
		TargetLang:  request.TargetLang,
	}
}
