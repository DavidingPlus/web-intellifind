package authentication

import (
	"backend/model/user"
	"backend/requests"
	"backend/tools/verifycode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SignUpService(c *gin.Context, signup_req requests.SignUpRequest) (err error) {

	exist := user.GetUser(signup_req.Email)
	//邮箱被注册
	if exist.Username != "" {
		return fmt.Errorf("该邮箱已被注册")
	}
	//校验验证码
	if flag := verifycode.NewVerifyCode().CheckAnswer(signup_req.Email, signup_req.VerifyCode); flag == false {
		//验证码错误
		return fmt.Errorf("验证码错误")
	}
	return nil

}
