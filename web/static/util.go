package static

import (
	"fmt"
	"math"
	"time"
)

func FormatRelativeTime(base time.Time, other time.Time) string {
	diff := base.Sub(other)

	if diff <= time.Hour {
		return fmt.Sprintf("%dm ago", int(math.Ceil(diff.Seconds()/60)))
	} else if diff < time.Hour {
		return fmt.Sprintf("%dh ago", int(diff.Minutes()/60))
	} else if diff < 24*time.Hour {
		return fmt.Sprintf("%dh ago", int(diff.Hours()))
	} else if diff < 30*24*time.Hour { // Up to 30 days
		return fmt.Sprintf("%dd ago", int(diff.Hours()/24))
	} else if diff < 365*24*time.Hour {
		return fmt.Sprintf("%dm ago", int(diff.Hours()/24/30)) // returns months
	} else {
		return fmt.Sprintf("%dy ago", int(diff.Hours()/24/365)) // returns years
	}
}
