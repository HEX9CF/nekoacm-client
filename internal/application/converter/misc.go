package converter

import (
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/pkg/pb"
)

// JokeToResponse 将笑话DTO转换为gRPC响应
func JokeToResponse(msg dto.ChatMsg) *pb.JokeResponse {
	return &pb.JokeResponse{
		Content: msg.Content,
	}
}

// ResponseToJoke 将gRPC响应转换为笑话DTO
func ResponseToJoke(resp *pb.JokeResponse) dto.ChatMsg {
	return dto.ChatMsg{
		Content: resp.Content,
	}
}

// JokeResponseToString 将笑话gRPC响应转换为字符串
func JokeResponseToString(resp *pb.JokeResponse) string {
	return resp.Content
}
