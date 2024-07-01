package helper

import (
	"strconv"
	"time"

	"github.com/polluxdev/trx-core-svc/common/utils"
)

func StringToTime(str, layout string) *time.Time {
	if str == "" {
		return nil
	}

	date, err := time.Parse(layout, str)
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	return &date
}

func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	return result
}

func TimeToString(date *time.Time, layout string) string {
	if date == nil {
		return ""
	}

	return date.Format(layout)
}
