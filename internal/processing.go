package internal

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sourabh-khot65/1brc/entity"
)

// ProcessData processes the data from the given file path and prints the results.
func ProcessData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("Error opening file : %v", err)
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		logrus.Errorf("Error getting file stat: %v", err)
		return err
	}
	fileSize := stat.Size()
	logrus.Infof("File size: %v", fileSize)

	cityMap := make(map[string]entity.CityMeasurements)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowData := strings.TrimSpace(scanner.Text())
		rowDataSplited := strings.Split(rowData, ";")
		if len(rowDataSplited) != 2 {
			logrus.Warnf("Skipping malformed line: %s", rowData)
			continue
		}

		cityName := strings.TrimSpace(rowDataSplited[0])
		temperatureStr := strings.TrimSpace(rowDataSplited[1])
		err = CalculateCityTemperatureMeasurements(cityMap, cityName, temperatureStr)
		if err != nil {
			logrus.Errorf("Error while calculating temp : %+v", err)
			continue
		}
	}

	// Print results
	for city, measurements := range cityMap {
		logrus.Infof("City: %s, Min: %s, Max: %s, Avg: %s, Count: %d", city, measurements.Min, measurements.Max, measurements.Avg, measurements.Count)
	}

	return nil
}
