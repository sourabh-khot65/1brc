package internal

import (
	"bufio"
	"os"

	"github.com/sirupsen/logrus"
)

func ProcessData(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Errorf("Error opening file : %v", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		logrus.Errorf("Error getting file stat: %v", err)
		return err
	}
	fileSize := stat.Size()
	buffer := make([]byte, fileSize)

	// Use bufio to read the file
	reader := bufio.NewReader(file)
	_, err = reader.Read(buffer)
	if err != nil {
		logrus.Errorf("Error reading file: %v", err)
		return err
	}
	logrus.Infof("File size: %v", fileSize)
	return nil
}
