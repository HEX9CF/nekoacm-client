package service

import (
	"context"
	"log"
	"nekoacm-client/internal/application/converter"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/infrastructure/neko"
	"strings"
	"time"
)

// 助手
func AssistantChat(msg dto.ChatMsg) (string, error) {
	msg.Content = strings.TrimSpace(msg.Content)
	if msg.Content == "" {
		return "", nil
	}
	log.Println("请求对话，内容长度:", len(msg.Content))

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.ChatMsgToRequest(msg)

	// 调用gRPC服务
	resp, err := neko.AssistantClient.Chat(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return "", err
	}

	// 使用converter转换PB到DTO
	chatResp := converter.ChatResponseToDto(resp)

	log.Println("请求成功，结果长度:", len(chatResp.Content))

	return chatResp.Content, nil
}
