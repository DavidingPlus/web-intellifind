package router

import (
	"backend/middlewares"
	"backend/service/authentication"
	"backend/service/cores"
	"backend/service/users"
	"github.com/gin-gonic/gin"
)

// 路由 api
func Router() *gin.Engine {
	r := gin.Default()

	//用户认证操作
	auth := r.Group("/auth", middlewares.GuestJWT())
	{

		auth.POST("/login", middlewares.LimitPerRoute("5-M"), authentication.LoginVerify)        //登陆
		auth.POST("/reset", authentication.Reset)                                                //重置密码
		auth.POST("/signup", authentication.SignUp)                                              //注册
		auth.POST("/send-code", middlewares.LimitPerRoute("1-M"), authentication.SendUsingEmail) //发送验证码
		auth.POST("verify-code", authentication.VerifyCode)                                      //校验验证码
		auth.POST("/capatcha-code", authentication.ShowCaptcha)                                  //获取图形验证码
		auth.POST("/is-exist", authentication.IsExist)                                           //邮箱是否注册

	}

	//用户相关
	user := r.Group("/user", middlewares.AuthJWT())
	{

		user.GET("/all", users.GetUserList) //获取所有用户信息，测试用
		//user.GET("/info", temp)      //获取用户信息
		//user.PUT("/update", temp)    //更新信息
		user.DELETE("/delete", users.DelteUser)         //删除用户
		user.POST("/refresh-token", users.RefreshToken) //刷新token
		user.POST("/upload-avatar", users.UploadAvatar) //上传头像

	}

	//核心功能
	core := r.Group("/core", middlewares.AuthJWT())
	{
		core.POST("/upload-file", middlewares.LimitPerRoute("1-M"), cores.UploadFile) //上传JSON文件
	}

	return r
}
