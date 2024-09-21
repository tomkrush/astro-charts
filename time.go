package main

import (
	"math"
	"time"
)

func JulianDate(date time.Time) float64 {
	year := date.Year()
	month := int(date.Month())
	day := date.Day()
	hour := date.Hour()
	minute := date.Minute()
	second := date.Second()

	if month <= 2 {
		year--
		month += 12
	}

	a := math.Floor(float64(year) / 100)
	b := 2 - a + math.Floor(a/4)

	julianDay := math.Floor(365.25*float64(year)) + math.Floor(30.6001*float64(month+1)) + float64(day) + 1720994.5 + b

	fractionOfDay := float64(hour)/24 + float64(minute)/1440 + float64(second)/86400

	julianDate := julianDay + fractionOfDay

	return julianDate
}
