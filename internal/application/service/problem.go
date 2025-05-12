package service

import (
	"context"
	"log"
	"nekoacm-client/internal/application/converter"
	"nekoacm-client/internal/application/dto"
	"nekoacm-client/internal/infrastructure/neko"
	"nekoacm-client/pkg/utils"
	"strconv"
	"time"
)

// 生成题目
func ProblemGenerate(pi dto.ProblemInstruction) (dto.Problem, error) {
	log.Println("请求生成题目...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.ProblemInstructionToRequest(pi)

	// 调用gRPC服务
	resp, err := neko.ProblemClient.Generate(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Problem{}, err
	}

	// 使用converter转换PB到DTO
	problem := converter.ResponseToProblem(resp)

	log.Println("生成题目成功")

	return problem, nil
}

// 解析题目
func ProblemParse(pd dto.ProblemData) (dto.Problem, error) {
	// 题目数据为空
	if pd.Data == "" {
		return dto.Problem{}, nil
	}

	log.Println("请求解析题目...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.ProblemDataToRequest(pd)

	// 调用gRPC服务
	resp, err := neko.ProblemClient.Parse(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Problem{}, err
	}

	// 使用converter转换PB到DTO
	problem := converter.ResponseToProblem(resp)

	log.Println("解析题目成功")

	return problem, nil
}

// 翻译题目
func ProblemTranslate(pi dto.TranslateInstruction) (dto.Problem, error) {
	log.Println("请求翻译题目...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.TranslateInstructionToRequest(pi)

	// 调用gRPC服务
	resp, err := neko.ProblemClient.Translate(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Problem{}, err
	}

	// 使用converter转换PB到DTO
	problem := converter.ResponseToProblem(resp)

	log.Println("翻译题目成功")

	return problem, nil
}

// 保存到json文件
func ProblemSaveJson(p dto.Problem) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/problem/" + p.Title + "_" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(p, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
