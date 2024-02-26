package tools

import (
	"github.com/spf13/viper"
	"time"
)

func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(viper.GetString("TimeZone"))
	return time.Now().In(chinaTimezone)
}
