package tests

import (
	"strconv"
	"testing"

	"github.com/sourabh-khot65/1brc/entity"
	"github.com/sourabh-khot65/1brc/internal"
)

func TestCalculateCityTemperatureMeasurements(t *testing.T) {
	cityMap := make(map[string]entity.CityMeasurements)

	tests := []struct {
		cityName      string
		temperature   string
		expectedMin   float64
		expectedMax   float64
		expectedAvg   float64
		expectedCount int64
		expectError   bool
	}{
		{"New York", "30", 30, 30, 30.00, 1, false},
		{"New York", "25", 25, 30, 27.50, 2, false},
		{"Los Angeles", "35", 35, 35, 35.00, 1, false},
		{"Chicago", "-10", -10, -10, -10.00, 1, false},
		{"Chicago", "invalid", 0, 0, 0, 0, true},
		{"Extreme City", "1000", 1000, 1000, 1000.00, 1, false},
		{"Extreme City", "-1000", -1000, 1000, -0.00, 2, false},
	}

	for _, test := range tests {
		tempValue, err := strconv.ParseFloat(test.temperature, 64)
		if err != nil {
			t.Errorf("Error parsing temperature for %s: %v", test.cityName, err)
		}

		err = internal.CalculateCityTemperatureMeasurements(cityMap, test.cityName, tempValue)

		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for city %s with temperature %s, but got none", test.cityName, test.temperature)
			}
			continue
		}

		if err != nil {
			t.Errorf("Error calculating temperature for %s: %v", test.cityName, err)
		}

		measurements := cityMap[test.cityName]
		if measurements.Min != test.expectedMin || measurements.Max != test.expectedMax || measurements.Avg != test.expectedAvg || measurements.Count != test.expectedCount {
			t.Errorf("For city %s, expected Min: %.2f, Max: %.2f, Avg: %.2f, Count: %d; got Min: %.2f, Max: %.2f, Avg: %.2f, Count: %d",
				test.cityName, test.expectedMin, test.expectedMax, test.expectedAvg, test.expectedCount,
				measurements.Min, measurements.Max, measurements.Avg, measurements.Count)
		}
	}
}
