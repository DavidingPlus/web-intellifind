package cores

import (
	"backend/requests"
	"backend/response"
	"backend/tools/file"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	req := requests.UploadJsonFileRequest{}
	c.ShouldBind(&req)
	// 上传文件
	save_path, file_name, err := file.SaveUploadJsonFile(c, req.JsonFile)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, "上传文件失败，请稍后尝试~")
		return
	}

	//解析json文件
	json_data, err := ParseJsonFile(c, save_path, file_name)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, "上传文件失败，请稍后尝试~")
		return
	}

	//解析成功后保存到数据库
	err = SaveJsonFile(json_data)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, "上传文件失败，请稍后尝试~")
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"message": "文件上传成功",
	})

}
