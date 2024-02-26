package authentication

import (
	"backend/tools/captcha"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseAPIController struct{}

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, _, err := captcha.NewCaptcha().GenerateCaptcha()
	if err != nil {
		fmt.Println(err)
	}

	// 返回给用户
	c.JSON(200, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
		//"answer":        answer,
	})
}
