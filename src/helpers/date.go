package helpers

import (
	"time"
)

func SubtractDaysInDate(date time.Time, days int) time.Time {
	return date.UTC().AddDate(0, 0, -1*days)
}
