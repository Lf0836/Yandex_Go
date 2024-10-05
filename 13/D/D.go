package main

import (
	"fmt"
	"time"

)

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	duration := now.Sub(pastTime)
	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%d seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	case duration < 30*24*time.Hour:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	case duration < 12*30*24*time.Hour:
		months := int(duration.Hours() / (24 * 30))
		return fmt.Sprintf("%d months ago", months)
	default:
		years := int(duration.Hours() / (24 * 365))
		return fmt.Sprintf("%d years ago", years)
	}
}
