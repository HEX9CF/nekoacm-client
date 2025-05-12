package bootstrap

import (
	"log"
	"nekoacm-client/internal/infrastructure/neko"
)

func initNekoAcm() error {
	log.Println("正在启动gRPC客户端...")
	err := neko.InitNekoAcm()
	if err != nil {
		log.Println("初始化gRPC客户端失败:", err)
		return err
	}

	return nil
}
