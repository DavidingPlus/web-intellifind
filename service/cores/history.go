package cores

import (
	"backend/model/core"
)

func GetHistory(uid uint, page int, size int) ([]core.Result, int, int, error) {
	history, err := core.GetHistory(uid)
	if err != nil {
		return nil, 0, 0, err
	}

	length := len(history)
	total_page := length/size + 1
	start := (page - 1) * size
	var end int
	if length > page*size {
		end = page * size
		return history[start:end], length, total_page, nil
	}
	return history[start:], length, total_page, nil

}

func GetResultOnce(fileName string) (core.SaveJsonFile, core.Settings, core.Result, error) {
	json_info, err := core.GetJsonInfo(fileName)
	settings, err := core.GetSettingByFileName(fileName)
	result, err := core.GetResultOnce(fileName)

	if err != nil {
		return core.SaveJsonFile{}, core.Settings{}, core.Result{}, err
	}
	return json_info, settings, result, nil
}

func DeleteHistoryOnceService(fileName string) error {
	err := core.DelteHistoryOnce(fileName)
	err = core.DeleteJsonInfo(fileName)
	err = core.DeleteSetting(fileName)
	if err != nil {
		return err
	}
	return nil
}

func ShowJsonInfoService(fileName string) (core.SaveJsonFile, error) {
	result, err := core.GetJsonInfo(fileName)
	if err != nil {
		return core.SaveJsonFile{}, err
	}
	return result, nil
}
