package config

type NekoAcmConf struct {
	Host    string `yaml:"host" json:"host"`
	Port    string `yaml:"port" json:"port"`
	Token   string `yaml:"token" json:"token"`
	Timeout int    `yaml:"timeout" json:"timeout"`
}

func (g *NekoAcmConf) Default() {
	g.Host = "localhost"
	g.Port = "14516"
	g.Token = "token"
	g.Timeout = 60
}
