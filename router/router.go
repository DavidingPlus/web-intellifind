package router

import (
	"backend/service/authentication"
	"backend/service/users"
	"github.com/gin-gonic/gin"
)

func temp(c *gin.Context) {
	c.JSON(200, gin.H{"data": "to do"})
}

func Router() *gin.Engine {
	r := gin.Default()

	//用户认证操作
	auth := r.Group("/auth")
	{
		vcc := new(authentication.VerifyCodeController)

		auth.POST("/login", authentication.LoginVerify) //登陆
		//auth.POST("/refresh-token", temp) //刷新token
		//auth.POST("/login", temp)         //重置密码
		auth.POST("/signup", authentication.SignUp)  //注册
		auth.POST("/send-code", vcc.SendUsingEmail)  //发送验证码
		auth.POST("/capatcha-code", vcc.ShowCaptcha) //获取图形验证码

	}

	////用户相关
	user := r.Group("/user")
	{

		user.GET("/all", users.GetUserList) //获取所有用户信息，测试用
		//user.GET("/info", temp)      //获取用户信息
		//user.PUT("/update", temp)    //更新信息
		user.DELETE("/delete", users.DelteUser) //删除用户

	}
	return r
}
