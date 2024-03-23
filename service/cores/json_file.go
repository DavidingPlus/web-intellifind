package cores

import (
	"backend/json"
	"backend/model/core"
	"github.com/gin-gonic/gin"
	"time"
)

func ParseJsonFile(c *gin.Context, file_path string, file_name string) (core.SaveJsonFile, error) {
	//todo 解析json文件

	json_parse, err := json.ParseJson(file_path)
	if err != nil {
		return core.SaveJsonFile{}, err
	}
	user_info := json_parse.Data[0]
	json_data := core.SaveJsonFile{
		FileName:   file_name,
		UID:        c.GetUint("current_user_id"),
		CreateTime: time.Now(),

		Desc:          user_info.Desc,
		Broswer:       user_info.EnvAttr.Browser.Value.(string),
		City:          user_info.EnvAttr.City.Value.(string),
		DeviceType:    user_info.EnvAttr.DeviceType.Value.(string),
		DisplayHeight: int(user_info.EnvAttr.DisplayHeight.Value.(float64)),
		DisplayWidth:  int(user_info.EnvAttr.DisplayWidth.Value.(float64)),

		Ip:              user_info.EnvAttr.Ip.Value.(string),
		OperatingSystem: user_info.EnvAttr.OperatingSystem.Value.(string),
		Province:        user_info.EnvAttr.Province.Value.(string),
		ProjectId:       user_info.ProjectId,
		ScreenDirect:    user_info.ScreenDirect,
		SessionId:       user_info.SessionId,
		UserId:          user_info.UserId,
	}

	return json_data, nil

}
func SaveJsonFile(json core.SaveJsonFile) error {
	//todo 保存json文件
	err := json.Create()
	return err

}
