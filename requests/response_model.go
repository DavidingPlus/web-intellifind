package requests

import "time"

type GetUserInfoResponse struct {
	UserName   string    `json:"user_name"`
	CreateTime time.Time `json:"create_time"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	City       string    `json:"city"`
	Gender     string    `gorm:"column:gender"`
	Tel        string    `gorm:"column:tel"`
	Birthday   string    `gorm:"column:birthday"`
}

type ShowHistoryResponse struct {
	CreateTime time.Time `json:"create_time"`
	FileName   string    `json:"file_name"`
	TotalScore float64   `json:"total_score"`

	PageError      float64 `json:"page_error"`
	PageLoad       float64 `json:"page_load"`
	PageExperience float64 `json:"page_experience"`
}

type ParseJsonResp struct {
	TotalScore float64 `json:"total_score,omitempty"` //总分

	//下面这一部分是每一个部分的得分
	StayTime         float64 `json:"stay_time,omitempty"`         // 1. 跳出率较高
	RepeatClick      float64 `json:"repeat_click,omitempty"`      // 2. 重复点击
	PageLoad         float64 ` json:"page_load,omitempty"`        //3.页面打开慢
	FeedbackInterval float64 `json:"feedback_interval,omitempty"` //4.点击后网络反馈慢
	NoReaction       float64 ` json:"no_reaction,omitempty"`      //5.点击无反应
	ErrorCount       float64 ` json:"error_count,omitempty"`      //6.点击报错
	ConsoleErrors    float64 `json:"console_errors,omitempty"`    //7.页面加载报错
	IsBlank          float64 ` json:"is_blank,omitempty"`         //8页面加载出现白屏
	OccurMany        float64 ` json:"occur_many,omitempty"`       //9.多个同时出现

	BriefDesc  string `json:"brief_desc,omitempty"`   //简单描述
	DetailDesc string ` json:"detail_desc,omitempty"` //详细描述
}
