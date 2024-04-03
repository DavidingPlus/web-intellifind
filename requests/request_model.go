package requests

import "mime/multipart"

type LoginRequest struct {
	Email     string `json:"email"  `
	Password  string `json:"password" `
	Answer    string `json:"answer"`
	CaptchaId string `json:"captcha_id"`
}

type SignUpRequest struct {
	Username   string `json:"username" `
	Email      string `json:"email" `
	Password   string `json:"password"  `
	VerifyCode string `json:"verify_code"`
}

type VerifyCaptchaRequest struct {
	Email string `json:"email"`
}

type VerifyCodeRequest struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

type ResetPasswordRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"`
}

type UpdatePasswordRequest struct {
	UID         string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type IsEmailExistRequest struct {
	Email string `json:"email"`
}

type UploadAvatarRequest struct {
	Avatar *multipart.FileHeader `form:"avatar"`
}
type UploadJsonFileRequest struct {
	JsonFile *multipart.FileHeader `form:"json_file"`

	//下面这一部分是每一个部分的得分
	StayTime         float64 ` json:"stay_time" form:"stay_time"`                 // 1. 跳出率较高
	RepeatClick      float64 `json:"repeat_click" form:"repeat_click"`            // 2. 重复点击
	PageLoad         float64 ` json:"page_load" form:"page_load"`                 //3.页面打开慢
	FeedbackInterval float64 ` json:"feedback_interval" form:"feedback_interval"` //4.点击后网络反馈慢
	NoReaction       float64 `json:"no_reaction" form:"no_reaction"`              //5.点击无反应
	ErrorCount       float64 ` json:"error_count" form:"error_count"`             //6.点击报错
	ConsoleErrors    float64 `json:"console_errors" form:"console_errors"`        //7.页面加载报错
	IsBlank          float64 ` json:"is_blank" form:"is_blank"`                   //8页面加载出现白屏
	OccurMany        float64 `json:"occur_many" form:"occur_many"`                //9.多个同时出现

}

type UpdateUserInfoRequest struct {
	UserName string `form:"user_name" json:"user_name"`
	City     string `form:"city" json:"city"`
	Gender   string `form:"column:gender" json:"gender"`
	Tel      string `form:"column:tel" json:"tel"`
	Birthday string `form:"column:birthday" json:"birthday"`
}

type ShowResultOnceRequest struct {
	FileName string `json:"file_name" form:"file_name"`
}

// 设置权重值
type EditSettingsRequest struct {
	StayTime         float64 `form:"stay_time"`         // 1. 跳出率较高
	RepeatClick      float64 `form:"repeat_click"`      // 2. 重复点击
	PageLoad         float64 `form:"page_load"`         //3.页面打开慢
	FeedbackInterval float64 `form:"feedback_interval"` //4.点击后网络反馈慢
	NoReaction       float64 `form:"no_reaction"`       //5.点击无反应
	ErrorCount       float64 `form:"error_count"`       //6.点击报错
	ConsoleErrors    float64 `form:"console_errors"`    //7.页面加载报错
	IsBlank          float64 `form:"is_blank"`          //8页面加载出现白屏
	OccurMany        float64 `form:"occur_many"`        //9.多个同时出现
}

// 分页展示历史记录
type ShowHistoryRequest struct {
	Size int `form:"size" json:"size"`
	Page int `form:"page" json:"page"`
}
