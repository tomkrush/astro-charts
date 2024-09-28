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
		st   ST
	}{
		{
			time.Date(2024, 9, 27, 23, 42, 56, 0, time.UTC),
			-88.35146,
			ST{
				GMSThh: 0,
				GMSTmm: 11,
				GMSTss: 56,
				LSThh:  18,
				LSTmm:  18,
				LSTss:  31,
			},
		},
		{
			time.Date(2000, 1, 1, 9, 30, 0, 0, time.UTC),
			-88.35146,
			ST{
				GMSThh: 16,
				GMSTmm: 11,
				GMSTss: 25,
				LSThh:  10,
				LSTmm:  18,
				LSTss:  1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.time.String(), func(t *testing.T) {
			st := SideRealTime(test.time, test.lon)

			if st.GMSThh != test.st.GMSThh {
				t.Errorf("Expected %v, got %v", test.st.GMSThh, st.GMSThh)
			}

			if st.GMSTmm != test.st.GMSTmm {
				t.Errorf("Expected %v, got %v", test.st.GMSTmm, st.GMSTmm)
			}

			if st.GMSTss != test.st.GMSTss {
				t.Errorf("Expected %v, got %v", test.st.GMSTss, st.GMSTss)
			}

			if st.LSThh != test.st.LSThh {
				t.Errorf("Expected %v, got %v", test.st.LSThh, st.LSThh)
			}

			if st.LSTmm != test.st.LSTmm {
				t.Errorf("Expected %v, got %v", test.st.LSTmm, st.LSTmm)
			}

			if st.LSTss != test.st.LSTss {
				t.Errorf("Expected %v, got %v", test.st.LSTss, st.LSTss)
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
