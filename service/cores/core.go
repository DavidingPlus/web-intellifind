package cores

import (
	"backend/model/core"
	"backend/requests"
	"backend/response"
	"backend/tools/file"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 上传json文件接口 同时会触发解析和保存
// todo 是否考虑把几个操作给分开
func UploadFile(c *gin.Context) {
	req := requests.UploadJsonFileRequest{}
	c.ShouldBind(&req)
	// 上传文件
	save_path, file_name, err := file.SaveUploadJsonFile(c, req.JsonFile)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, err.Error())
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

	//获取解析的结果
	uid := c.GetUint("current_user_id")
	settings, _ := GetSettingService(uid)
	data, err := GetResult(c, save_path, file_name, settings)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code":    1,
		"message": "解析成功",
		"data":    data,
	})

}

func ShowHistory(c *gin.Context) {
	current_uid := c.GetUint("current_user_id")

	history, err := GetHistory(current_uid)

	if err != nil {
		response.Abort500(c, "查询历史失败")
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": history,
	})
}
func ShowResultOnce(c *gin.Context) {

	req := requests.ShowResultOnceRequest{}
	c.ShouldBind(&req)
	data, err := GetResultOnce(req.FileName)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": data,
	})

}

func GetSetting(c *gin.Context) {
	uid := c.GetUint("current_user_id")
	settings, err := GetSettingService(uid)
	if err != nil {
		if settings.UID == 0 {
			default_setting, _ := CreateDefaultSettings(uid)
			c.JSON(200, gin.H{
				"code":    1,
				"message": "默认设置",
				"data":    default_setting,
			})
			return
		}
		response.Abort500(c, "获取设置失败")
		return
	}
	c.JSON(200, gin.H{
		"code":    1,
		"message": "用户设置如下",
		"data":    settings,
	})

}

func EditSettings(c *gin.Context) {
	req := core.Settings{}
	c.ShouldBind(&req)

	uid := c.GetUint("current_user_id")
	req.UID = uid
	err := EditSetting(req)

	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": "更新成功",
	})

}
