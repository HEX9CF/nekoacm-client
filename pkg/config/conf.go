package config

import (
	"nekoacm-common/pkg/config"
	"nekoacm-common/pkg/utils"
)

var (
	Conf *Config = &Config{}
)

type Config struct {
	Server  ServerConf       `yaml:"server" json:"server"`
	NekoAcm NekoAcmConf      `yaml:"nekoacm" json:"nekoacm"`
	Nacos   config.NacosConf `yaml:"nacos" json:"nacos"`
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
	c.NekoAcm.Default()
	c.Nacos.Default()
}
