package main

import (
	"backend/router"
	"backend/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.ShowConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()
	r.Run(viper.GetString("port"))

}
