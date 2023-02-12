package config

type ServerConfig struct {
	Name             string           `mapstructure:"name" json:"name"`
	Host             string           `mapstructure:"host" json:"host"`
	Port             int64            `mapstructure:"port" json:"port"`
	Tags             []string         `mapstructure:"tags" json:"tags"`
	PayServerConfig  PayServerConfig  `mapstructure:"pay-server" json:"pay-server"`
	UserServerConfig UserServerConfig `mapstructure:"user-server" json:"user-server"`
	AliPayConfig     AliPayConfig     `mapstructure:"alipay" json:"alipay"`
}

type AliPayConfig struct {
	AppID        string `mapstructure:"app_id" json:"app_id"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyUrl    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnUrl    string `mapstructure:"return_url" json:"return_url"`
	IsProduction bool   `mapstructure:"is_production" json:"is_production"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type PayServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
