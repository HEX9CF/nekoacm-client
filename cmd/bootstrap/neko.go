package bootstrap

import (
	"log"
	"nekoacm-client/internal/infrastructure/neko"
)

func initNekoAcm() error {
	err := neko.InitNekoAcm()
	if err != nil {
		log.Println("初始化 gRPC 客户端失败！")
		return err
	}
	log.Println("初始化 gRPC 客户端成功")

	return nil
}
