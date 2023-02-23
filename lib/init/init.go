package internal

import (
	"gateway/lib/config"
	"gateway/lib/logger"
	"gateway/lib/mysql"
	"go.uber.org/zap"
	"os"
)

func InitModule(configPath string) {
	// 初始化日志记录
	logger.InitLogger()

	if configPath == "" {
		zap.S().Info("input config file like config.yml")
		os.Exit(1)
	}
	// 读取配置文件
	config.InitConfig(configPath)

	// 初始化数据库句柄
	mysql.InitMysql()
}
