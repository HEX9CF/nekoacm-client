package bootstrap

import (
	"log"
	"os"
)

// 初始化
func Init() error {
	log.SetOutput(os.Stdout)

	if err := initConfig(); err != nil {
		return err
	}

	if err := initNacos(); err != nil {
		return nil
	}

	if err := initNekoAcm(); err != nil {
		return err
	}

	return nil
}
