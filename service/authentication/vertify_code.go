package authentication

import (
	"backend/requests"
	"backend/tools/verifycode"
	"github.com/gin-gonic/gin"
)

func SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.VerifyCaptchaRequest{}
	c.ShouldBind(&request)

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

func VerifyCode(c *gin.Context) {
	req := requests.VerifyCodeRequest{}
	if flag := verifycode.NewVerifyCode().CheckAnswer(req.Email, req.VerifyCode); flag == false {
		//验证码错误
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误！",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    1,
		"message": "验证成功",
	})
}
