package model

import (
	"backend/model/core"
	"backend/model/user"
)

//MySQL表名

type tales struct {
	UserInfos            user.UserInfos            `gorm:"user_infos"`
	UploadJsonFileReocrd core.UploadJsonFileRecord `gorm:"upload_json_file_record"`
	SaveJsonFile         core.SaveJsonFile         `gorm:"save_json_file"`
}
