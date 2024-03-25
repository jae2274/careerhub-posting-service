package utils

import "time"

func UnixMilliToTime(unixMilli int64) time.Time {
	seconds := unixMilli / 1000
	nanoseconds := (unixMilli % 1000) * 1e6
	return time.Unix(seconds, nanoseconds)
}

func UnixMilliToTimePtr(unixMilli *int64) *time.Time {
	if unixMilli == nil {
		return nil
	}
	result := UnixMilliToTime(*unixMilli)
	return &result
}
