package service

import (
	"context"
	"log"
	"nekoacm-client/internal/application/converter"
	"nekoacm-client/internal/infrastructure/neko"
	"nekoacm-common/api/proto/pb"
	"time"
)

// 生成笑话
func TellJoke() (string, error) {
	log.Println("请求生成笑话...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 调用gRPC服务
	resp, err := neko.MiscClient.TellJoke(ctx, &pb.Empty{})

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return "", err
	}

	// 使用converter转换PB到字符串
	joke := converter.JokeResponseToString(resp)

	log.Println("生成笑话成功")

	return joke, nil
}
