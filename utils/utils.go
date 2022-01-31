package utils

import "time"

func Time(hour, min, sec int) time.Duration {
	return time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
}
