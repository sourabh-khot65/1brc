package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/sourabh-khot65/1brc/entity"
)

// ProcessData processes the data from the given file path and prints the results.
func ProcessData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := stat.Size()
	logrus.Infof("File size: %v", fileSize)

	cityMap := make(map[string]entity.CityMeasurements)
	var mu sync.Mutex
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		wg.Add(1)
		go func(rowData string) {
			defer wg.Done()
			rowData = strings.TrimSpace(rowData)
			rowDataSplited := strings.Split(rowData, ";")
			if len(rowDataSplited) != 2 {
				logrus.Warnf("Skipping malformed line: %s", rowData)
				return
			}

			cityName := strings.TrimSpace(rowDataSplited[0])
			temperatureStr := strings.TrimSpace(rowDataSplited[1])
			temperature, err := strconv.ParseFloat(temperatureStr, 64)
			if err != nil {
				logrus.Errorf("Invalid temperature value: %s", temperatureStr)
				return
			}

			mu.Lock()
			defer mu.Unlock()
			err = CalculateCityTemperatureMeasurements(cityMap, cityName, temperature)
			if err != nil {
				logrus.Errorf("Error while calculating temp: %+v", err)
			}
		}(scanner.Text())
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		return err
	}

	// Print results
	for city, measurements := range cityMap {
		logrus.Infof("City: %s, Min: %.2f, Max: %.2f, Avg: %.2f, Count: %d", city, measurements.Min, measurements.Max, measurements.Avg, measurements.Count)
	}

	return nil
}
