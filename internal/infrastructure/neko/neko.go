package neko

import "nekoacm-client/pkg/config"

func InitNekoAcm() error {
	var err error
	conf := config.Conf.NekoAcm
	address := conf.Host + ":" + conf.Port

	err = createConn(address)
	if err != nil {
		return err
	}

	initClient()

	return nil
}
