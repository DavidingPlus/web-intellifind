package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

// 展示配置初始化信息
func ShowConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf("config for mysql,", viper.GetString("MySql.dsn"))
	fmt.Sprintf("config for redis,", viper.GetString("Redis.dsn"))
}
