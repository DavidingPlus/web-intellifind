package users

import (
	"backend/model"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*model.UserInfos, 10)
	data = model.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

func DelteUser(c *gin.Context) {
	id := c.Query("id")
	model.DelteUser(id)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}
