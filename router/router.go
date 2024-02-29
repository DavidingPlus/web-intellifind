package router

import (
	"backend/middlewares"
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

		auth.POST("/login", middlewares.GuestJWT(), middlewares.LimitPerRoute("5-M"), authentication.LoginVerify)        //登陆
		auth.POST("/reset", middlewares.GuestJWT(), authentication.Reset)                                                //重置密码
		auth.POST("/signup", middlewares.GuestJWT(), authentication.SignUp)                                              //注册
		auth.POST("/send-code", middlewares.GuestJWT(), middlewares.LimitPerRoute("1-M"), authentication.SendUsingEmail) //发送验证码
		auth.POST("verify-code", middlewares.GuestJWT(), authentication.VerifyCode)                                      //校验验证码
		auth.POST("/capatcha-code", middlewares.GuestJWT(), authentication.ShowCaptcha)                                  //获取图形验证码

	}

	////用户相关
	user := r.Group("/user")
	{

		user.GET("/all", users.GetUserList) //获取所有用户信息，测试用
		//user.GET("/info", temp)      //获取用户信息
		//user.PUT("/update", temp)    //更新信息
		user.DELETE("/delete", middlewares.AuthJWT(), users.DelteUser) //删除用户
		//user.PUT("/update-password", users.UpdatePassword) //修改密码 考虑一下要不要做这个

		user.POST("/refresh-token", middlewares.AuthJWT(), users.RefreshToken) //刷新token

	}
	return r
}
