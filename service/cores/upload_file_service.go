package cores

import (
	"backend/model/core"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetResult(c *gin.Context, save_path string, file_name string) (core.Result, error) {
	url_value := url.Values{"file_name": {save_path}}

	//调用解析接口

	//todo 在算法部分起一个服务
	resp, err := http.PostForm("http://127.0.0.1:8080/api/parse", url_value)
	if err != nil {
		return core.Result{}, fmt.Errorf("请求接口失败")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var res core.Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Errorf("解析失败")
	}

	result := core.Result{
		FileName:         file_name,
		UID:              c.GetUint("current_user_id"),
		CreateTime:       time.Now(),
		TotalScore:       res.TotalScore,
		StayTime:         res.StayTime,
		RepeatClick:      res.RepeatClick,
		PageLoad:         res.PageLoad,
		FeedbackInterval: res.FeedbackInterval,
		NoReaction:       res.NoReaction,
		ErrorCount:       res.ErrorCount,
		ConsoleErrors:    res.ConsoleErrors,
		IsBlank:          res.IsBlank,
		OccurMany:        res.OccurMany,
		BriefDesc:        res.BriefDesc,
		DetailDesc:       res.DetailDesc,
	}
	err = result.Create()
	if err != nil {
		return core.Result{}, fmt.Errorf("保存失败")
	}
	return result, nil
}
