package authentication

import (
	"backend/jwt"
	"backend/model/user"
	"backend/requests"
	"backend/tools/captcha"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LoginService(c *gin.Context, login_request requests.LoginRequest) (token string, err error) {
	if match := captcha.NewCaptcha().VerifyCaptcha(login_request.CaptchaId, login_request.Answer); match == false {
		return "", fmt.Errorf("验证码错误")
	}

	//校验用户名密码
	flag := user.ComparePassword(login_request.Password, login_request.Email)
	//登陆失败：用户名密码错误
	if !flag {
		return "", fmt.Errorf("用户名或密码错误")
	}

	//生成token
	data := user.GetUser(login_request.Email)

	ID_string := strconv.Itoa(int(data.ID))
	token = jwt.NewJWT().IssueToken(ID_string, data.Username)

	//登陆成功
	return token, nil

}
