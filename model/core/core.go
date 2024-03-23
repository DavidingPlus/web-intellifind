package core

import (
	"backend/utils"
	"fmt"
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
func GetFileNamebyUID(uid uint) ([]string, error) {
	var filenames []string

	result := utils.DB.Select("file_name").Where("uid = ?", uid).Find(&filenames)
	if result.RowsAffected == 0 {
		return filenames, fmt.Errorf("查询为空")
	}
	return filenames, nil

}
