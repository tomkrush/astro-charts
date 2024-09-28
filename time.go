package main

import (
	"math"
	"time"
)

type ST struct {
	GMSThh int
	GMSTmm int
	GMSTss int
	LSThh  int
	LSTmm  int
	LSTss  int
}

type SiderealTimeResult struct {
	LocalSiderealTime     LocalSiderealTime
	GreenwichSiderealTime GreenwichSiderealTime
}

type LocalSiderealTime struct {
	hh int
	mm int
	ss int
}

type GreenwichSiderealTime struct {
	hh int
	mm int
	ss int
}

func SideRealTime(date time.Time, lon float64) SiderealTimeResult {
	jd := JulianDate(date)

	LongDeg := math.Abs(lon)
	LongMin := (LongDeg - math.Floor(LongDeg)) * 60
	LongSec := (LongMin - math.Floor(LongMin)) * 60
	LongMin = math.Floor(LongMin)
	LongSec = math.Floor(LongSec)

	GMST := 18.697374558 + 24.06570982441908*(jd-2451545.0)
	GMST = math.Mod(GMST, 24)
	GMSTmm := (GMST - math.Floor(GMST)) * 60
	GMSTss := (GMSTmm - math.Floor(GMSTmm)) * 60
	GMSThh := math.Floor(GMST)
	GMSTmm = math.Floor(GMSTmm)
	GMSTss = math.Floor(GMSTss)

	Long := lon / 15
	LST := GMST + Long
	if LST < 0 {
		LST += 24
	}

	LSTmm := (LST - math.Floor(LST)) * 60
	LSTss := (LSTmm - math.Floor(LSTmm)) * 60
	LSThh := math.Floor(LST)
	LSTmm = math.Floor(LSTmm)
	LSTss = math.Floor(LSTss)

	return SiderealTimeResult{
		LocalSiderealTime: LocalSiderealTime{
			hh: int(LSThh),
			mm: int(LSTmm),
			ss: int(LSTss),
		},
		GreenwichSiderealTime: GreenwichSiderealTime{
			hh: int(GMSThh),
			mm: int(GMSTmm),
			ss: int(GMSTss),
		},
	}
}

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
