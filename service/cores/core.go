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

	//进行解析
	uid := c.GetUint("current_user_id")
	settings := core.Settings{
		UID:              uid,
		FileName:         file_name,
		StayTime:         req.StayTime,
		RepeatClick:      req.RepeatClick,
		PageLoad:         req.PageLoad,
		FeedbackInterval: req.FeedbackInterval,
		NoReaction:       req.NoReaction,
		ErrorCount:       req.ErrorCount,
		ConsoleErrors:    req.ConsoleErrors,
		IsBlank:          req.IsBlank,
		OccurMany:        req.OccurMany,
	}
	core.CreateSetting(settings)

	_, err = GetResult(c, save_path, file_name, settings)
	if err != nil {
		fmt.Println(err)
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code":      1,
		"message":   "解析成功",
		"file_name": file_name,
	})

}

func ShowHistory(c *gin.Context) {

	req := requests.ShowHistoryRequest{}
	c.ShouldBind(&req)

	current_uid := c.GetUint("current_user_id")

	data, length, total_page, err := GetHistory(current_uid, req.Page, req.Size)

	if err != nil {
		response.Abort500(c, "查询历史失败")
		return
	}
	c.JSON(200, gin.H{
		"code":       1,
		"length":     length,
		"total_page": total_page,
		"data":       data,
	})
}
func ShowResultOnce(c *gin.Context) {

	req := requests.ShowResultOnceRequest{}
	c.ShouldBind(&req)
	json_info, settings, res, err := GetResultOnce(req.FileName)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	data := make(map[string]interface{})
	data["basic_info"] = json_info
	data["settings"] = settings
	data["result"] = res

	c.JSON(200, gin.H{
		"code": 1,
		"data": data,
	})

}

func GetSettingLast(c *gin.Context) {
	uid := c.GetUint("current_user_id")
	settings, err := GetSettingLastService(uid)
	if err != nil {
		if settings.UID == 0 {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "默认设置",
				"data":    settings,
			})
			return
		}
		response.Abort500(c, "获取设置失败")
		return
	}
	c.JSON(200, gin.H{
		"code":    1,
		"message": "用户上一次解析的设置如下",
		"data":    settings,
	})

}

//func EditSettings(c *gin.Context) {
//	req := core.Settings{}
//	c.ShouldBind(&req)
//
//	uid := c.GetUint("current_user_id")
//	req.UID = uid
//	err := EditSetting(req)
//
//	if err != nil {
//		response.Abort500(c, err.Error())
//		return
//	}
//	c.JSON(200, gin.H{
//		"code": 1,
//		"data": "更新成功",
//	})
//
//}

func DeleteHistoryOnce(c *gin.Context) {

	req := requests.ShowResultOnceRequest{}
	c.ShouldBind(&req)
	err := DeleteHistoryOnceService(req.FileName)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": "删除成功",
	})
}

func ShowJsonInfo(c *gin.Context) {
	req := requests.ShowResultOnceRequest{}
	c.ShouldBind(&req)

	res, err := ShowJsonInfoService(req.FileName)

	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"data": res,
	})

}
