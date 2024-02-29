package requests

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
