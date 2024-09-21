package main

import (
	"math"
	"testing"
	"time"
)

func Float64Equal(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
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
