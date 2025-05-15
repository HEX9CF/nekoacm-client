package bootstrap

import (
	"log"
	"nekoacm-client/internal/infrastructure/nacos"
	"nekoacm-client/internal/interfaces/http"
	"nekoacm-client/pkg/config"
)

// 初始化服务器
func InitServer() error {
	if config.Conf.Nacos.Enable {
		if err := nacos.NacosClient.Register(); err != nil {
			log.Println("注册服务到 Nacos 失败！")
			return err
		}
	}

	// 启动HTTP服务器
	log.Println("正在启动 HTTP 服务器...")
	if err := http.InitServer(); err != nil {
		log.Println("初始化 HTTP 服务器失败！")
		return err
	}

	return nil
}
