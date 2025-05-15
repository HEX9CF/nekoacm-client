package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-common/api/proto/pb"
)

// TestcaseInstructionToRequest 将测试用例指令DTO转换为gRPC请求
func TestcaseInstructionToRequest(instruction dto.TestcaseInstruction) *pb.TestcaseInstructionRequest {
	return &pb.TestcaseInstructionRequest{
		Title:        instruction.Title,
		Description:  instruction.Description,
		Input:        instruction.Input,
		Output:       instruction.Output,
		SampleInput:  instruction.SampleInput,
		SampleOutput: instruction.SampleOutput,
		Hint:         instruction.Hint,
		Tags:         instruction.Tags,
		Solution:     instruction.Solution,
	}
}

// RequestToTestcaseInstruction 将gRPC请求转换为测试用例指令DTO
func RequestToTestcaseInstruction(request *pb.TestcaseInstructionRequest) dto.TestcaseInstruction {
	return dto.TestcaseInstruction{
		Title:        request.Title,
		Description:  request.Description,
		Input:        request.Input,
		Output:       request.Output,
		SampleInput:  request.SampleInput,
		SampleOutput: request.SampleOutput,
		Hint:         request.Hint,
		Tags:         request.Tags,
		Solution:     request.Solution,
	}
}

// TestcaseToResponse 将测试用例DTO转换为gRPC测试用例
func TestcaseToResponse(testcase dto.Testcase) *pb.TestcaseResponse {
	return &pb.TestcaseResponse{
		TestInput:         testcase.TestInput,
		TestOutput:        testcase.TestOutput,
		InputExplanation:  testcase.InputExplanation,
		OutputExplanation: testcase.OutputExplanation,
	}
}

// ResponseToTestcase 将gRPC测试用例转换为DTO测试用例
func ResponseToTestcase(tc *pb.TestcaseResponse) dto.Testcase {
	return dto.Testcase{
		TestInput:         tc.TestInput,
		TestOutput:        tc.TestOutput,
		InputExplanation:  tc.InputExplanation,
		OutputExplanation: tc.OutputExplanation,
	}
}
