package nacos

import (
	"nekoacm-client/pkg/config"
	"nekoacm-common/infrastructure/nacos"
)

var (
	NacosClient *nacos.NacosClient
)

// InitNacos 初始化Nacos服务
func InitNacos() error {
	conf := &config.Conf.Nacos

	// 使用server包中的配置创建nacos客户端
	NacosClient = nacos.NewNacosClient(conf)

	// 初始化nacos客户端
	if err := NacosClient.Init(); err != nil {
		return err
	}

	// 注册服务
	if conf.Register.Enable {
		if err := NacosClient.Register(); err != nil {
			return err
		}
	}

	return nil
}
