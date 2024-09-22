package main

import (
	"math"
	"testing"
	"time"
)

func Float64Equal(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestSiderealTime(t *testing.T) {
	tests := []struct {
		time time.Time
		lon  float64
		st   float64
	}{
		{
			time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC),
			-71.0589,
			6.6461,
		},
		{
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			-71.0589,
			6.6461,
		},
	}

	for _, test := range tests {
		t.Run(test.time.String(), func(t *testing.T) {
			st := SideRealTime(test.time, test.lon)
			if !Float64Equal(st, test.st, 0.0001) {
				t.Errorf("Expected %v, got %v", test.st, st)
			}
		})
	}
}

func TestJulianDate(t *testing.T) {
	tests := []struct {
		time time.Time
		jd   float64
	}{
		{
			time.Date(2024, 9, 21, 0, 0, 0, 0, time.UTC),
			2460574.5,
		},
		{
			time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			2451544.5,
		},
	}

	for _, test := range tests {
		t.Run(test.time.String(), func(t *testing.T) {
			jd := JulianDate(test.time)
			if !Float64Equal(jd, test.jd, 0.0001) {
				t.Errorf("Expected %v, got %v", test.jd, jd)
			}
		})
	}
}
