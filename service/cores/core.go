package cores

import (
	"backend/requests"
	"backend/response"
	"backend/tools/file"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	req := requests.UploadJsonFileRequest{}
	c.ShouldBind(&req)
	//todo 上传文件
	_, err := file.SaveUploadJsonFile(c, req.JsonFile)
	if err != nil {
		response.Abort500(c, "上传文件失败，请稍后尝试~")
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "文件上传成功",
	})

}
