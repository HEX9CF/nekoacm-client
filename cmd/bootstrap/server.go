package bootstrap

import (
	"log"
	"nekoacm-client/internal/interfaces/http"
)

// 初始化服务器
func InitServer() error {
	// 启动HTTP服务器
	log.Println("正在启动HTTP服务器...")
	err := http.InitServer()
	if err != nil {
		log.Println("初始化HTTP服务器失败！")
		return err
	}

	return nil
}
