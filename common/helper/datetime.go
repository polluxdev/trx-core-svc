package helper

import "time"

func GetNow(duration time.Duration) *time.Time {
	now := time.Now().Add(duration)
	return &now
}
