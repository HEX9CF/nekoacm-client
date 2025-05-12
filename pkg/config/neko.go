package config

type NekoGrpcConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (g *NekoGrpcConf) Default() {
	g.Host = "localhost"
	g.Port = "14516"
	g.Token = "token"
}
