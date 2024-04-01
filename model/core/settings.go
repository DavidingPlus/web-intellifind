package core

import (
	"backend/utils"
	"fmt"
)

type Settings struct {
	FileName string `gorm:"primary_key" json:"file_name"`
	UID      uint   `gorm:"uid" json:"uid"`

	//下面这一部分是每一个部分的得分
	StayTime         float64 `gorm:"stay_time" json:"stay_time" form:"stay_time"`                         // 1. 跳出率较高
	RepeatClick      float64 `gorm:"repeat_click" json:"repeat_click" form:"repeat_click"`                // 2. 重复点击
	PageLoad         float64 `gorm:"page_load" json:"page_load" form:"page_load"`                         //3.页面打开慢
	FeedbackInterval float64 `gorm:"feedback_interval" json:"feedback_interval" form:"feedback_interval"` //4.点击后网络反馈慢
	NoReaction       float64 `gorm:"no_reaction" json:"no_reaction" form:"no_reaction"`                   //5.点击无反应
	ErrorCount       float64 `gorm:"error_count" json:"error_count" form:"error_count"`                   //6.点击报错
	ConsoleErrors    float64 `gorm:"console_errors" json:"console_errors" form:"console_errors"`          //7.页面加载报错
	IsBlank          float64 `gorm:"is_blank" json:"is_blank" form:"is_blank"`                            //8页面加载出现白屏
	OccurMany        float64 `gorm:"occur_many" json:"occur_many" form:"occur_many"`                      //9.多个同时出现

}

func (s *Settings) TableName() string {
	return "settings"
}

func CreateSetting(req Settings) error {

	result := utils.DB.Create(&req)

	if result.RowsAffected == 0 {
		return fmt.Errorf("创建失败")
	}
	return nil
}

func DeleteSetting(file_name string) error {

	result := utils.DB.Where("file_name = ?", file_name).Delete(&Settings{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("删除失败")
	}
	return nil

}

func GetSettingByUid(uid uint) (Settings, error) {
	var res Settings
	result := utils.DB.Where("uid = ?", uid).Last(&res)
	if result.RowsAffected == 0 {
		return Settings{}, fmt.Errorf("查询为空")
	}
	return res, nil
}

func GetSettingByFileName(fileName string) (Settings, error) {
	var res Settings
	result := utils.DB.Where("file_name = ?", fileName).First(&res)
	if result.RowsAffected == 0 {
		return Settings{}, fmt.Errorf("查询为空")
	}
	return res, nil
}

//func EditSettings(req Settings) error {
//	var err error
//	err = DeleteSetting(req.UID)
//	err = CreateSetting(req)
//	if err != nil {
//		return fmt.Errorf("更新失败")
//	}
//	return nil
//}
