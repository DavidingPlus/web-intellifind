package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonParse struct {
	Data    []Info `json:"data"`
	Message string `json:"message"`
}
type Info struct {
	Desc            string      `json:"desc"`
	DeviceGroupId   int         `json:"deviceGroupId"`
	EnvAttr         EnvAttr     `json:"envAttr"`
	EventId         string      `json:"eventId"`
	EventName       string      `json:"eventName"`
	EventTypeId     int         `json:"eventTypeId"`
	InteractionAttr interface{} `json:"interactionAttr"`
	PageAttr        interface{} `json:"pageAttr"`
	PerformanceAttr interface{} `json:"performanceAttr"`
	ProjectId       int         `json:"projectId"`
	ScreenDirect    int         `json:"screenDirect"`
	SessionId       int         `json:"sessionId"`
	UserId          int         `json:"userId"`
}
type EnvAttr struct {
	Browser         Detail `json:"browser"`
	City            Detail `json:"city"`
	Country         Detail `json:"country"`
	DeviceType      Detail `json:"deviceType"`
	DisplayHeight   Detail `json:"displayHeight"`
	DisplayWidth    Detail `json:"displayWidth"`
	Ip              Detail `json:"ip"`
	OperatingSystem Detail `json:"operatingSystem"`
	Province        Detail `json:"province"`
}
type Detail struct {
	DisplayName string
	Order       string
	Value       interface{}
}

func ParseJson(file_path string) (JsonParse, error) {
	file, _ := os.ReadFile(file_path)
	json_parse := JsonParse{}
	json.Unmarshal(file, &json_parse)
	if len(json_parse.Data) == 0 {
		err := error(fmt.Errorf("上传文件为空或未解析到data字段"))
		return json_parse, err
	}

	return json_parse, nil

}
