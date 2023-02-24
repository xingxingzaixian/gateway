package mysql

import (
	"fmt"
	"gateway/lib/config"
	"gateway/models"
	"gateway/public"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
		})
	mysqlConfig := config.BaseConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		zap.S().Fatalf("数据库连接失败：%s:%d", mysqlConfig.Host, mysqlConfig.Port)
	}

	db.AutoMigrate(&models.HttpRule{}, &models.ServiceInfo{})
	public.GormDB = db
}
