package config

type Socket struct {
	Start bool   `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Addr  string `mapstructure:"addr" json:"addr" yaml:"addr"`    // socket监听地址:端口
}
