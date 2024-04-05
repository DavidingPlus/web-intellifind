package router

import (
	"backend/middlewares"
	"backend/service/authentication"
	"backend/service/cores"
	"backend/service/users"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"

	"strings"
)

// 路由 api
func Router() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})

	r.Use(middlewares.Cors())

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

		user.GET("/all", users.GetUserList)             //获取所有用户信息，测试用
		user.GET("/info", users.GetUserInfo)            //获取用户信息
		user.PUT("/update", users.UpdateUserInfo)       //更新信息
		user.DELETE("/delete", users.DelteUser)         //删除用户
		user.POST("/refresh-token", users.RefreshToken) //刷新token
		user.POST("/upload-avatar", users.UploadAvatar) //上传头像
		user.GET("/logout", users.Logout)               //用户登出

	}

	//核心功能
	core := r.Group("/core", middlewares.AuthJWT())
	{
		core.POST("/upload-file", middlewares.LimitPerRoute("2-M"), cores.UploadFile) //上传JSON文件
		//core.PUT("/settings/edit_my", cores.EditSettings)                             // 编辑配置
		core.GET("/settings/get", cores.GetSettingLast)              // 获得用户上一次解析的配置
		core.GET("/show-history/total", cores.ShowHistory)           //解析的历史记录
		core.GET("/show-result/once", cores.ShowResultOnce)          //查看具体的某一次解析
		core.DELETE("/delete-history/once", cores.DeleteHistoryOnce) //删除某一次历史记录
		core.GET("/show/json-info", cores.ShowJsonInfo)              //展示用户上传的json文件中的基本信息
	}

	return r
}
