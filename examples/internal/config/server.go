package config

var ServerConfig = new(map[string]Server)

type Server struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"` // 监听地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"` // 监听端口
	ReadTimeout  int    `json:"read-timeout" yaml:"read-timeout"`     // 读超时
	WriteTimeout int    `json:"write-timeout" yaml:"write-timeout"`   // 写超时
}
