package utils

import (
	"backend/logger"
	"github.com/spf13/viper"
)

func InitLogger() {
	logger.InitLogger(
		viper.GetString("Log.filename"),
		viper.GetInt("Log.max_size"),
		viper.GetInt("Log.max_backup"),
		viper.GetInt("Log.max_age"),
		viper.GetBool("Log.compress"),
		viper.GetString("Log.type"),
		viper.GetString("Log.level"),
	)
}
