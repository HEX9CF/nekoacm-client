package nacos

import "nekoacm-client/pkg/config"

func LoadConfig() error {
	var err error

	if err = NacosClient.GetConfig(&config.Conf.Server, "nekoacm-client-server.yaml"); err != nil {
		return err
	}

	if err = NacosClient.GetConfig(&config.Conf.NekoAcm, "nekoacm-client-nekoacm.yaml"); err != nil {
		return err
	}

	return nil
}
