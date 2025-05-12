package service

import (
	"context"
	"log"
	"nekoacm-client/internal/application/converter"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/infrastructure/neko"
	"time"
)

// 提交评测
func Submit(s dto.Submission) (dto.Judgement, error) {
	log.Println("请求评测代码...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.SubmissionToRequest(s)

	// 调用gRPC服务
	resp, err := neko.JudgeClient.Submit(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Judgement{}, err
	}

	log.Println("评测代码成功")

	// 使用converter转换PB到DTO
	judgement := converter.ResponseToJudgement(resp)

	return judgement, nil
}
