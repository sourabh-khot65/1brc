package internal

import (
	"strconv"

	"github.com/sourabh-khot65/1brc/entity"
)

// CalculateCityTemperatureMeasurements calculates the temperature measurements for a given city.
func CalculateCityTemperatureMeasurements(cityMap map[string]entity.CityMeasurements, cityName, temperature string) error {
	tempValue, err := strconv.ParseFloat(temperature, 64)
	if err != nil {
		return err // Return the error if parsing fails
	}

	prevData, exist := cityMap[cityName]
	if !exist {
		cityMap[cityName] = entity.CityMeasurements{
			Min:   temperature,
			Max:   temperature,
			Avg:   temperature,
			Count: 1,
		}
	} else {
		minValue, _ := strconv.ParseFloat(prevData.Min, 64)
		maxValue, _ := strconv.ParseFloat(prevData.Max, 64)
		newCount := prevData.Count + 1

		// Update Min and Max
		if tempValue < minValue {
			prevData.Min = temperature
		}
		if tempValue > maxValue {
			prevData.Max = temperature
		}

		// Calculate new average
		totalTemperature := (minValue*float64(prevData.Count) + tempValue) / float64(newCount)
		prevData.Avg = strconv.FormatFloat(totalTemperature, 'f', 2, 64)
		prevData.Count = newCount

		cityMap[cityName] = prevData
	}
	return nil // Return nil if everything is successful
}
