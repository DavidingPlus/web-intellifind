package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	//全局的DB
	DB *gorm.DB
)

// mysql初始化
func InitMysql() {
	newlogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢sql阈值
		LogLevel:      logger.Info, //级别
		Colorful:      true,        //彩色
	},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("MySql.dsn")), &gorm.Config{Logger: newlogger})
	if err != nil {
		panic("failed to connect database")
	}
}
