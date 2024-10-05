package main

import "time"

func FormatTimeToString(time time.Time, format string) string {
	return time.Format(format)
}
