package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SwaggerInfo struct {
	Title    string `mapstructure:"title"`
	Desc     string `mapstructure:"desc"`
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
}

type WebInfo struct {
	Url            string `mapstructure:"url"`
	Path           string `mapstructure:"path"`
	Addr           string `mapstructure:"addr"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

type MysqlInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type ProxyInfo struct {
	Addr           string `mapstructure:"addr"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

type HttpInfo struct {
	Addr           string `mapstructure:"addr"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

type HttpsInfo struct {
	Addr           string `mapstructure:"addr"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

type Config struct {
	DebugMode    string      `mapstructure:"debug_mode"`
	TimeLocation string      `mapstructure:"time_location"`
	Salt         string      `mapstructure:"salt"`
	JwtSignKey   string      `mapstructure:"sign-key"`
	Swagger      SwaggerInfo `mapstructure:"swagger"`
	Mysql        MysqlInfo   `mapstructure:"mysql"`
	ProxyHttp    ProxyInfo   `mapstructure:"proxy.http"`
	ProxyHttps   ProxyInfo   `mapstructure:"proxy.https"`
	Web          WebInfo     `mapstructure:"web"`
	Http         HttpInfo    `mapstructure:"http"`
}

var BaseConfig Config

func InitConfig(configPath string) {
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("读取配置文件【%s】失败", configPath)
	}

	if err := v.Unmarshal(&BaseConfig); err != nil {
		zap.S().Panicf("配置文件【%s】格式异常", configPath)
	}
	zap.S().Infof("配置文件【%s】读取成功", configPath)
	zap.S().Infof("配置信息：%v", BaseConfig)
}
