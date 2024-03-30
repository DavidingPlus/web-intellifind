package core

import (
	"backend/utils"
	"fmt"
	"time"
)

type SaveJsonFile struct {
	FileName   string    `gorm:"primary_key"`
	UID        uint      `gorm:"uid"`
	CreateTime time.Time `gorm:"column:create_time"`

	Desc            string
	Broswer         string
	City            string
	DeviceType      string
	DisplayHeight   int
	DisplayWidth    int
	Ip              string
	OperatingSystem string
	Province        string
	ProjectId       int
	ScreenDirect    int
	SessionId       int
	UserId          int
}

func (s *SaveJsonFile) Create() error {
	res := utils.DB.Create(&s)
	if res.RowsAffected == 0 {
		return error(fmt.Errorf("创建失败"))
	}
	return nil
}

func GetJsonInfo(fileName string) (SaveJsonFile, error) {
	var res SaveJsonFile
	result := utils.DB.Where("file_name = ?", fileName).Find(&res)
	if result.RowsAffected == 0 {
		return SaveJsonFile{}, fmt.Errorf("查询为空")
	}
	return res, nil
}
func DeleteJsonInfo(fileName string) error {
	result := utils.DB.Where("file_name = ?", fileName).Delete(&SaveJsonFile{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("删除失败")
	}
	return nil
}
