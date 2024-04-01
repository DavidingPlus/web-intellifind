package cores

import (
	"backend/model/core"
)

var default_settings = core.Settings{
	FileName:         "default_settings",
	StayTime:         0.15,
	RepeatClick:      0.05,
	PageLoad:         0.15,
	FeedbackInterval: 0.15,
	NoReaction:       0.1,
	ErrorCount:       0.1,
	ConsoleErrors:    0.1,
	IsBlank:          0.05,
	OccurMany:        0.15,
}

func GetSettingLastService(uid uint) (core.Settings, error) {
	setting, err := core.GetSettingByUid(uid)

	if err != nil {
		return default_settings, err
	}
	return setting, err

}

//
//func EditSetting(req core.Settings) error {
//
//	err := core.EditSettings(req)
//	if err != nil {
//		return err
//	}
//	return nil
//
//}
