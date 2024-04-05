package core

import (
	"backend/utils"
	"fmt"
	"time"
)

type Result struct {
	FileName   string    `gorm:"primary_key" json:"file_name"`                    //文件名字
	UID        uint      `gorm:"uid" json:"uid,omitempty"`                        //用户id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time,omitempty"` //时间
	TotalScore float64   `gorm:"total_score" json:"total_score,omitempty"`        //总分

	//下面这一部分是每一个部分的得分
	StayTime         float64 `gorm:"stay_time" json:"stay_time,omitempty"`                 // 1. 跳出率较高
	RepeatClick      float64 `gorm:"repeat_click" json:"repeat_click,omitempty"`           // 2. 重复点击
	PageLoad         float64 `gorm:"page_load" json:"page_load,omitempty"`                 //3.页面打开慢
	FeedbackInterval float64 `gorm:"feedback_interval" json:"feedback_interval,omitempty"` //4.点击后网络反馈慢
	NoReaction       float64 `gorm:"no_reaction" json:"no_reaction,omitempty"`             //5.点击无反应
	ErrorCount       float64 `gorm:"error_count" json:"error_count,omitempty"`             //6.点击报错
	ConsoleErrors    float64 `gorm:"console_errors" json:"console_errors,omitempty"`       //7.页面加载报错
	IsBlank          float64 `gorm:"is_blank" json:"is_blank,omitempty"`                   //8页面加载出现白屏
	OccurMany        float64 `gorm:"occur_many" json:"occur_many,omitempty"`               //9.多个同时出现

	BriefDesc  string `gorm:"brief_desc" json:"brief_desc,omitempty"`    //简单描述
	DetailDesc string `gorom:"detail_desc" json:"detail_desc,omitempty"` //详细描述
}

func (r *Result) TableName() string {
	return "result"
}

func (r *Result) Create() error {
	res := utils.DB.Create(r)
	if res.RowsAffected == 0 {
		return error(fmt.Errorf("创建失败"))
	}
	return nil

}

func GetHistory(uid uint) ([]Result, error) {
	var history []Result

	result := utils.DB.Where("uid = ?", uid).Order("create_time desc").Find(&history)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("查询为空")
	}
	return history, nil

}

func GetResultOnce(fileName string) (Result, error) {
	var res Result
	result := utils.DB.Where("file_name = ?", fileName).First(&res)

	if result.RowsAffected == 0 {
		return Result{}, fmt.Errorf("查询为空")
	}
	return res, nil
}

func DelteHistoryOnce(fileName string) error {
	result := utils.DB.Where("file_name = ?", fileName).Delete(&Result{})
	if result.RowsAffected == 0 {
		return fmt.Errorf("删除失败")
	}
	return nil
}
