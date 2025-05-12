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

// 生成测试用例
func TestcaseGenerate(ti dto.TestcaseInstruction) (dto.Testcase, error) {
	log.Println("请求生成测试用例...")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 使用converter转换DTO到PB
	req := converter.TestcaseInstructionToRequest(ti)

	// 调用gRPC服务
	resp, err := neko.TestcaseClient.Generate(ctx, req)

	if err != nil {
		log.Println("gRPC调用失败:", err)
		return dto.Testcase{}, err
	}

	// 使用converter转换PB到DTO
	testcase := converter.ResponseToTestcase(resp)

	log.Println("生成测试用例成功")

	return testcase, nil
}

// 保存到json文件
func TestcaseSaveJson(t dto.Testcase) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/testcase/" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(t, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
