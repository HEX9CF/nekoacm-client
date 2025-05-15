package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-common/api/proto/pb"
)

// ChatMsgToRequest 将DTO转换为gRPC请求
func ChatMsgToRequest(msg dto.ChatMsg) *pb.ChatRequest {
	return &pb.ChatRequest{
		Content: msg.Content,
	}
}

// ChatResponseToDto 将gRPC响应转换为DTO
func ChatResponseToDto(resp *pb.ChatResponse) dto.ChatMsg {
	return dto.ChatMsg{
		Content: resp.Content,
	}
}
