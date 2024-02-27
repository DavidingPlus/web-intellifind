package authentication

import (
	"backend/jwt"
	"backend/model/user"
	"backend/requests"
	"backend/tools/captcha"
	"backend/tools/verifycode"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	//校验用户名密码
	flag := user.ComparePassword(login_request.Password, login_request.Email)

	//登陆失败：用户名密码错误
	if !flag {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
		return
	}

	//生成token
	data := user.GetUserByEmail(login_request.Email)

	ID_string := strconv.Itoa(int(data.ID))
	token := jwt.NewJWT().IssueToken(ID_string, data.Username)

	//登陆成功
	c.JSON(200, gin.H{
		"code":    1,
		"message": "登陆成功",
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
	exist := user.GetUserByEmail(signup_req.Email)

	//邮箱被注册
	if exist.Username != "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "邮箱已被注册",
		})
		return
	}

	//校验验证码
	if flag := verifycode.NewVerifyCode().CheckAnswer(signup_req.Email, signup_req.VerifyCode); flag == false {
		//验证码错误
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "验证码错误！",
		})
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
