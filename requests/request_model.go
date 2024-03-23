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
	CaptchaId string `json:"captcha_id"`
	Answer    string `json:"answer"`
	Email     string `json:"email"`
}

type VerifyCodeRequest struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verify_code"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
}

type UpdateUserInfoRequest struct {
	UserName string `form:"user_name"`
	City     string `form:"city"`
}

type ShowResultOnceRequest struct {
	FileName string `json:"file_name"`
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
