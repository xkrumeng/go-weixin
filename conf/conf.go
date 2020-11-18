package conf

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// App 配置
	App appConfig

	// Server 服务器配置
	Server serverConfig

	// Wx 微信配置
	Wx weixinConfig

	// Mysql 配置
	Mysql mysqlConfig
)

type appConfig struct {
	RunMode string `ini:"run_mode"`
}

type serverConfig struct {
	HTTPPort     string        `ini:"http_port"`
	ReadTimeout  time.Duration `ini:"read_timeout"`
	WriteTimeout time.Duration `ini:"write_timeout"`
}

type weixinConfig struct {
	AppID     string `ini:"appId"`
	AppSecret string `ini:"appSecret"`
	Token     string `ini:"token"`
}

type mysqlConfig struct {
	Type        string `ini:"type"`
	User        string `ini:"user"`
	Password    string `ini:"password"`
	Host        string `ini:"host"`
	Name        string `ini:"name"`
	TablePrefix string `ini:"table_prefix"`
}

// LoadConfig 加载配置文件
func LoadConfig() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	err = cfg.Section("app").MapTo(&App)
	if err != nil {
		log.Fatal("Failed to load App Config")
	}

	err = cfg.Section("server").MapTo(&Server)
	if err != nil {
		log.Fatal("Failed to load Server Config")
	}

	err = cfg.Section("weixin").MapTo(&Wx)
	if err != nil {
		log.Fatal("Failed to load weixin Config")
	}

	err = cfg.Section("mysql").MapTo(&Mysql)
	if err != nil {
		log.Fatal("Failed to load mysql Config")
	}
}
