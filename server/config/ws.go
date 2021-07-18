package config

type WebSocket struct {
	Start bool   `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Port  string `mapstructure:"port" json:"port" yaml:"port"`    // websocket监听地址:端口
}
