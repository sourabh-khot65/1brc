package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sourabh-khot65/1brc/internal"
)

func TestProcessData(t *testing.T) {
	tests := []struct {
		fileContent string
		expectError bool
		expectedLog []string
	}{
		{
			fileContent: "New York;30\nLos Angeles;25\nNew York;35\nChicago;20\nLos Angeles;28\n",
			expectError: false,
			expectedLog: []string{
				"City: New York, Min: 30.00, Max: 35.00, Avg: 32.50, Count: 2",
				"City: Los Angeles, Min: 25.00, Max: 28.00, Avg: 26.50, Count: 2",
				"City: Chicago, Min: 20.00, Max: 20.00, Avg: 20.00, Count: 1",
			},
		},
		{
			fileContent: "New York;invalid\nLos Angeles;25\n",
			expectError: true,
			expectedLog: []string{
				"Invalid temperature value: invalid",
			},
		},
		{
			fileContent: "MalformedLineWithoutSemicolon\nLos Angeles;25\n",
			expectError: false,
			expectedLog: []string{
				"Skipping malformed line: MalformedLineWithoutSemicolon",
				"City: Los Angeles, Min: 25.00, Max: 25.00, Avg: 25.00, Count: 1",
			},
		},
		{
			fileContent: "Extreme City;1000\nExtreme City;-1000\n",
			expectError: false,
			expectedLog: []string{
				"City: Extreme City, Min: -1000.00, Max: 1000.00, Avg: 0.00, Count: 2",
			},
		},
	}

	for _, test := range tests {
		// Create a temporary file with the test content
		tmpFile, err := os.CreateTemp("", "testdata")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(test.fileContent); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Close()

		// Capture log output
		var logOutput []string
		logrus.SetOutput(&logWriter{&logOutput})
		defer logrus.SetOutput(os.Stderr)

		// Run the ProcessData function
		err = internal.ProcessData(tmpFile.Name())
		if test.expectError {
			if err == nil {
				t.Errorf("Expected error but got none")
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}

		// Check log output
		for _, expectedLog := range test.expectedLog {
			found := false
			for _, log := range logOutput {
				if strings.Contains(log, expectedLog) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected log output: %s, but not found", expectedLog)
			}
		}
	}
}

// logWriter is a helper type to capture log output
type logWriter struct {
	logs *[]string
}

func (w *logWriter) Write(p []byte) (n int, err error) {
	*w.logs = append(*w.logs, string(p))
	return len(p), nil
}
