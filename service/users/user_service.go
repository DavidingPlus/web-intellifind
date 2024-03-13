package users

import (
	"backend/jwt"
	"backend/logger"
	"backend/model/user"
	"backend/requests"
	"backend/response"
	"backend/tools/file"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.UserInfos {
	userModel, ok := c.MustGet("current_user").(user.UserInfos)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.UserInfos{}
	}
	// db is now a *DB value
	return userModel
}

func GetUserList(c *gin.Context) {
	data := make([]*user.UserInfos, 10)
	data = user.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

func DelteUser(c *gin.Context) {
	id, _ := c.Get("current_user_id")
	id_string := fmt.Sprintf("%v", id)

	user.DelteUser(id_string)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// RefreshToken 刷新 Access Token
func RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"code":  1,
			"token": token,
		})
	}
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {

	id, _ := c.Get("current_user_id")
	id_string := fmt.Sprintf("%v", id)

	//todo 展示哪些信息 具体还是等user表定下来
	userInfo := user.GetUser(id_string)

	c.JSON(200, gin.H{
		"code": 1,
		"data": userInfo,
	})

}

// 上传头像
func UploadAvatar(c *gin.Context) {

	req := requests.UploadAvatarRequest{}
	c.ShouldBind(&req)

	avatar, err := file.SaveUploadAvatar(c, req.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := CurrentUser(c)
	currentUser.Avatar = avatar
	currentUser.Save()

	c.JSON(200, gin.H{
		"code":    1,
		"message": "头像上传成功",
	})

}
