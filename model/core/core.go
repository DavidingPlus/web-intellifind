package core

import (
	"backend/utils"
	"time"
)

type UploadJsonFileRecord struct {
	FileName   string    `gorm:"primary_key"`
	UID        uint      `gorm:"column:uid"`
	UploadTime time.Time `gorm:"column:upload_time"`
	SavePath   string    `gorm:"column:save_path"`
}

func (u *UploadJsonFileRecord) TableName() string {
	return "upload_json_file_record"
}
func (u *UploadJsonFileRecord) Create() {
	utils.DB.Create(&u)

}
