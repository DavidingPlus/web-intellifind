package cores

import (
	"backend/model/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/HttpRequest"
	"github.com/spf13/viper"
	"time"
)

func GetResult(c *gin.Context, save_path string, file_name string, settings core.Settings) (core.Result, error) {

	req := HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{"Content-Type": "application/json"})
	data := map[string]interface{}{
		"save_path":      save_path,
		"stay_time":      settings.StayTime,
		"repeat_click":   settings.RepeatClick,
		"page_load":      settings.PageLoad,
		"feedback":       settings.FeedbackInterval,
		"no_reaction":    settings.NoReaction,
		"error_count":    settings.ErrorCount,
		"console_errors": settings.ConsoleErrors,
		"is_blank":       settings.IsBlank,
		"occur_many":     settings.OccurMany,
	}

	url := viper.GetString("Algorithm.url")
	res, err := req.Post(url, data)

	if err != nil {
		return core.Result{}, err
	}

	result := core.Result{}
	res.Json(&result)

	result.UID = settings.UID
	result.FileName = file_name
	result.CreateTime = time.Now()

	fmt.Println(result)

	err = result.Create()
	if err != nil {
		return core.Result{}, fmt.Errorf("保存失败")
	}
	return result, nil
}
