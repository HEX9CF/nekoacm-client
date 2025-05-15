package service

import (
	"context"
	"log"
	"nekoacm-client/internal/application/converter"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/infrastructure/neko"
	"nekoacm-common/pkg/utils"
	"strconv"
	"time"
)

// 生成题解
func SolutionGenerate(si dto.SolutionInstruction) (dto.Solution, error) {
	log.Println("请求生成题解...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.SolutionInstructionToRequest(si)

	// 调用gRPC服务
	resp, err := neko.SolutionClient.Generate(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Solution{}, err
	}

	log.Println("生成题解成功，代码长度:", len(resp.SourceCode))

	// 使用converter转换PB到DTO
	solution := converter.ResponseToSolution(resp)

	return solution, nil
}

// 保存到json文件
func SolutionSaveJson(s dto.Solution) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/solution/" + s.Language + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(s, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
