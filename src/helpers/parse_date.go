package helpers

import "time"

func ParseDate(dateStr string) time.Time {
	dateFormat := "2006/01/02"

	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {

		return time.Now()
	}

	return date
}

func IsFutureDate(date time.Time) bool {
	now := time.Now()
	return !date.Before(now)
}