package middlewares

import (
	"backend/jwt"
	"backend/model/user"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()

	}
}

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请先登陆再进行操作"))
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.GetUser(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.ID)
		c.Set("current_user_name", userModel.Username)
		c.Set("current_user", userModel)

		c.Next()
	}
}
