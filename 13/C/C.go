package main

import "time"

func ParseStringToTime(data, format string) (time.Time, error) {
	return time.Parse(format, data)
}
