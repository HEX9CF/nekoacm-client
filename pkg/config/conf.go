package config

import (
	"nekoacm-client/pkg/utils"
)

type Config struct {
	Server   ServerConf   `yaml:"server" json:"server"`
	NekoGrpc NekoGrpcConf `yaml:"neko-grpc" json:"neko-grpc"`
}

// InitConfig 初始化
func InitConfig() error {
	v, err := utils.IsFileExists("config.yaml")
	if err != nil {
		return err
	}
	if !v {
		Conf.Default()
		err = utils.WriteYaml(&Conf, "config.yaml")
		if err != nil {
			return err
		}
	}
	err = utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Default() {
	c.Server.Default()
	c.NekoGrpc.Default()
}
