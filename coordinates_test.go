package main

import (
	"testing"
)

func TestRAToHours(t *testing.T) {
	// testing table
	tests := []struct {
		ra    string
		hours float64
	}{
		{
			ra:    "18h 44m 20.4s",
			hours: 18.739,
		},
		{
			ra:    "05h 55m 10.3s",
			hours: 5.919,
		},
	}

	for _, test := range tests {
		t.Run(test.ra, func(t *testing.T) {
			hours := ParseHours(test.ra)

			if !Float64Equal(hours, test.hours, 0.001) {
				t.Errorf("Expected %v, got %v", test.hours, hours)
			}
		})
	}
}

func TestDECToDegrees(t *testing.T) {
	// testing table
	tests := []struct {
		dec     string
		degrees float64
	}{
		{
			dec:     "+38° 47′ 01″",
			degrees: 38.7836,
		},
	}

	for _, test := range tests {
		t.Run(test.dec, func(t *testing.T) {
			degrees := ParseDegrees(test.dec)

			if !Float64Equal(degrees, test.degrees, 0.001) {
				t.Errorf("Expected %v, got %v", test.degrees, degrees)
			}
		})
	}
}
