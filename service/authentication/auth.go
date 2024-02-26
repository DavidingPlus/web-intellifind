package authentication

import (
	"backend/jwt"
	"backend/model"
	"backend/requests"
	"backend/tools/captcha"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 登录校验

func LoginVerify(c *gin.Context) {

	login_request := requests.LoginRequest{}

	err := c.ShouldBind(&login_request)
	if err != nil {
		fmt.Println(err)
	}

	//校验图形验证码
	if match := captcha.NewCaptcha().VerifyCaptcha(login_request.CaptchaId, login_request.Answer); match == false {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}

	data := model.UserInfos{}
	utils.DB.Where("email = ? AND password = ?", login_request.Email, login_request.Password).First(&data)

	if data.Username == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
		return
	}

	ID_string := strconv.Itoa(int(data.ID))
	token := jwt.NewJWT().IssueToken(ID_string, data.Username)

	c.JSON(200, gin.H{
		"code":    1,
		"message": "登陆成功",
		"token":   token,
	})

}

func SignUp(c *gin.Context) {

	signup_req := requests.SignUpRequest{}
	err := c.ShouldBind(&signup_req)

	if err != nil {
		fmt.Println(err)
	}
	exist := model.GetUserByEmail(signup_req.Email)
	if exist.Username != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "邮箱已被注册",
		})
		return
	}

	//todo 发送验证码

	//todo 校验验证码

	flag := true

	if flag == true {

		data := model.UserInfos{
			Username: signup_req.Username,
			Email:    signup_req.Email,
			Password: signup_req.Password,
		}

		model.CreateUser(data)

		c.JSON(200, gin.H{
			"code":    1,
			"message": "注册成功",
			"data":    data,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误！",
		})
	}

}
