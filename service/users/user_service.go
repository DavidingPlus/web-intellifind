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

// 获取用户列表 测试用
func GetUserList(c *gin.Context) {
	data := make([]*user.UserInfos, 10)
	data = user.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// 删除用户 测试用
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

	token, err := jwt.NewJWT().RefreshToken(c, true)

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

	res := requests.GetUserInfoResponse{
		UserName:   userInfo.Username,
		Email:      userInfo.Email,
		CreateTime: userInfo.CreatedAt,
		Avatar:     userInfo.Avatar,
		City:       userInfo.City,
		Gender:     userInfo.Gender,
		Tel:        userInfo.Tel,
		Birthday:   userInfo.Birthday,
	}

	c.JSON(200, gin.H{
		"code": 1,
		"data": res,
	})
}

// 上传头像
func UploadAvatar(c *gin.Context) {

	req := requests.UploadAvatarRequest{}
	c.ShouldBind(&req)

	avatar, err := file.SaveUploadAvatar(c, req.Avatar)
	if err != nil {
		response.Abort500(c, err.Error())
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

// 更新其他信息（用户名，城市）
func UpdateUserInfo(c *gin.Context) {

	req := requests.UpdateUserInfoRequest{}
	c.ShouldBind(&req)

	currentUser := CurrentUser(c)
	currentUser.Username = req.UserName
	currentUser.City = req.City
	currentUser.Tel = req.Tel
	currentUser.Gender = req.Gender
	currentUser.Birthday = req.Birthday

	err := currentUser.Save()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "更新成功",
	})
}

func Logout(c *gin.Context) {
	token, err := jwt.NewJWT().SetInvalidToken(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "登出成功！",
		"token":   token,
	})

}
