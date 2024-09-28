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
		st   SiderealResult
	}{
		{
			time.Date(2024, 9, 27, 23, 42, 56, 0, time.UTC),
			-88.35146,
			SiderealResult{
				LocalSiderealTime: SiderealTime{
					hh: 18,
					mm: 18,
					ss: 31,
				},
				GreenwichSiderealTime: SiderealTime{
					hh: 0,
					mm: 11,
					ss: 56,
				},
			},
		},
		{
			time.Date(2000, 1, 1, 9, 30, 0, 0, time.UTC),
			-88.35146,
			SiderealResult{
				LocalSiderealTime: SiderealTime{
					hh: 10,
					mm: 18,
					ss: 1,
				},
				GreenwichSiderealTime: SiderealTime{
					hh: 16,
					mm: 11,
					ss: 25,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.time.String(), func(t *testing.T) {
			st := ConvertToST(test.time, test.lon)

			if st.LocalSiderealTime != test.st.LocalSiderealTime {
				t.Errorf("Expected %v, got %v", test.st.LocalSiderealTime, st.LocalSiderealTime)
			}

			if st.GreenwichSiderealTime != test.st.GreenwichSiderealTime {
				t.Errorf("Expected %v, got %v", test.st.GreenwichSiderealTime, st.GreenwichSiderealTime)
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
