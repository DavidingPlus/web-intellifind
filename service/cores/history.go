package cores

import (
	"backend/model/core"
)

func GetHistory(uid uint) ([]core.Result, error) {
	history, err := core.GetHistory(uid)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func GetResultOnce(fileName string) (core.Result, error) {
	result, err := core.GetResultOnce(fileName)
	if err != nil {
		return core.Result{}, err
	}
	return result, nil
}
