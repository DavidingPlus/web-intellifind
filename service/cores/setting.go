package cores

import "backend/model/core"

func GetSettingService(uid uint) (core.Settings, error) {
	setting, err := core.GetSetting(uid)

	if err != nil {
		return core.Settings{}, err
	}
	return setting, err

}

func EditSetting(req core.Settings) error {

	err := core.EditSettings(req)
	if err != nil {
		return err
	}
	return nil

}

func CreateDefaultSettings(uid uint) (core.Settings, error) {
	default_settings := core.Settings{
		UID:              uid,
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
	err := core.CreateSetting(default_settings)
	if err != nil {
		return core.Settings{}, err
	}
	return default_settings, nil
}
