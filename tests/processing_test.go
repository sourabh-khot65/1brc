package tests

import (
	"os"
	"testing"

	"github.com/sourabh-khot65/1brc/internal"
)

func TestProcessData(t *testing.T) {
	tests := []struct {
		fileContent string
		expectError bool
	}{
		{
			fileContent: "New York;30\nLos Angeles;25\nNew York;35\nChicago;20\nLos Angeles;28\n",
			expectError: false,
		},
		{
			fileContent: "New York;invalid\nLos Angeles;25\n",
			expectError: true,
		},
		{
			fileContent: "MalformedLineWithoutSemicolon\nLos Angeles;25\n",
			expectError: false,
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
	}
}
