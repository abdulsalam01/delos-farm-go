package helper

import (
	"time"
)

var (
	timeNow = time.Now
)

func ConvertInt64ToTime(datetime int64) time.Time {
	return time.Unix(datetime, 0)
}

func ConvertTimeToInt64(datetime time.Time) int64 {
	return datetime.Unix()
}

func ConvertStringToTime(datetime string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", datetime)
	if err != nil {
		return date, err
	}

	return date, nil
}

func GetTimeNow() time.Time {
	return timeNow()
}
