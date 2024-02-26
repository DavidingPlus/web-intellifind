package authentication

import (
	"backend/requests"
	"backend/tools/captcha"
	"github.com/gin-gonic/gin"
)

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	c.ShouldBind(&request)
	if match := captcha.NewCaptcha().VerifyCaptcha(request.CaptchaId, request.Answer); match != true {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}

	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "发送失败",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "发送成功",
		})
	}
}
