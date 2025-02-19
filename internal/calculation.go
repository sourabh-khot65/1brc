package internal

import (
	"github.com/sourabh-khot65/1brc/entity"
)

// CalculateCityTemperatureMeasurements calculates the temperature measurements for a given city.
func CalculateCityTemperatureMeasurements(cityMap map[string]entity.CityMeasurements, cityName string, temperature float64) error {
	prevData, exist := cityMap[cityName]
	if !exist {
		cityMap[cityName] = entity.CityMeasurements{
			Min:   temperature,
			Max:   temperature,
			Avg:   temperature,
			Count: 1,
		}
	} else {
		minValue := prevData.Min
		maxValue := prevData.Min
		newCount := prevData.Count + 1

		// Update Min and Max
		if temperature < minValue {
			prevData.Min = temperature
		}
		if temperature > maxValue {
			prevData.Max = temperature
		}

		// Calculate new average
		totalTemperature := (minValue*float64(prevData.Count) + temperature) / float64(newCount)
		prevData.Avg = totalTemperature
		prevData.Count = newCount

		cityMap[cityName] = prevData
	}
	return nil // Return nil if everything is successful
}
