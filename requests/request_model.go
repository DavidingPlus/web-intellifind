package requests

type LoginRequest struct {
	Email     string `json:"email"  `
	Password  string `json:"password" `
	Answer    string `json:"answer"`
	CaptchaId string `json:"captcha_id"`
}

type SignUpRequest struct {
	Username string `json:"username" `
	Email    string `json:"email" `
	Password string `json:"password"  `
}

type VerifyCodeEmailRequest struct {
	CaptchaId string `json:"captcha_id"`
	Answer    string `json:"answer"`
	Email     string `json:"email"`
}
