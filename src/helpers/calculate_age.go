package helpers

import (
	"time"
)

func CalculateAge(bornDate time.Time) int {
	currentTime := time.Now()

	age := currentTime.Year() - bornDate.Year()

	if currentTime.Month() < bornDate.Month() ||
		(currentTime.Month() == bornDate.Month() && currentTime.Day() < bornDate.Day()) {
		age--
	}

	return age
}
