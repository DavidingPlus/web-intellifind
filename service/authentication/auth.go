package authentication

import (
	"backend/model/user"
	"backend/requests"
	"backend/tools/verifycode"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 登录校验
func LoginVerify(c *gin.Context) {

	login_request := requests.LoginRequest{}

	err := c.ShouldBind(&login_request)
	if err != nil {
		fmt.Println(err)
	}

	token, err := LoginService(c, login_request)

	if err != nil {
		c.JSON(401, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "登录成功",
		"token":   token,
	})

}

// 注册
func SignUp(c *gin.Context) {

	signup_req := requests.SignUpRequest{}
	err := c.ShouldBind(&signup_req)

	if err != nil {
		fmt.Println(err)
	}

	if err := SignUpService(c, signup_req); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	} else {
		data := user.UserInfos{
			Username: signup_req.Username,
			Email:    signup_req.Email,
			Password: signup_req.Password,
		}

		user.CreateUser(data)

		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册成功",
			"data":    data,
		})
	}

}

// 重置密码
func Reset(c *gin.Context) {

	req := requests.ResetPasswordRequest{}

	c.ShouldBind(&req)

	if flag := verifycode.NewVerifyCode().CheckAnswer(req.Email, req.VerifyCode); flag == false {
		//验证码错误
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误！",
		})
		return
	}

	user.UpdatePassword(req.Password, req.Email)

	c.JSON(200, gin.H{
		"code":    1,
		"message": "重置成功",
	})

}

func IsExist(c *gin.Context) {

	req := requests.IsEmailExistRequest{}
	c.ShouldBind(&req)
	flag := user.GetUser(req.Email)

	if flag.Username == "" {
		c.JSON(200, gin.H{
			"code":    2,
			"message": "邮箱未注册",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "邮箱已注册",
		})
	}
}
