package core

import (
	"backend/utils"
	"time"
)

type UploadJsonFileRecord struct {
	UID        uint
	UploadTime time.Time
	FileName   string
	SavePath   string
}

func (u *UploadJsonFileRecord) TableName() string {
	return "upload_json_file_record"
}
func (u *UploadJsonFileRecord) Create() {
	utils.DB.Create(&u)

}
