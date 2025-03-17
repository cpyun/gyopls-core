package config

var ApplicationConfig = new(Application)

type Application struct {
	Mode     string `json:"mode" yaml:"mode"`           //环境配置 dev开发环境 test测试环境 prod线上环境
	Name     string `json:"name" yaml:"name"`           // 服务名称
	EnableDp bool   `json:"enable-dp" yaml:"enable-dp"` // 数据权限功能开关
}
