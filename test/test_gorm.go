package main

import (
	"backend/model/core"
	"backend/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:qz1731978669@tcp(8.137.100.0:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&user.UserInfos{})
	db.AutoMigrate(&core.UploadJsonFileRecord{})
	db.AutoMigrate(&core.SaveJsonFile{})
	db.AutoMigrate(&core.Settings{})
	db.AutoMigrate(&core.Result{})

}
