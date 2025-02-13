package main

import (
	"log"
	"os"

	"github.com/sourabh-khot65/1brc/internal"
	"github.com/sourabh-khot65/1brc/profiling"
)

func main() {
	// Start CPU profiling and defer its stop.
	// The StartCPUProfiling function is expected to return the profileFile and an error.
	profileFile, err := profiling.StartCPUProfiling()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		profiling.StopCPUProfiling()
		profileFile.Close()
	}()
	if len(os.Args) != 2 {
		log.Fatalf("Missing measurements filename")
	}
	fileName := os.Args[1]
	err = internal.ProcessData(fileName)
	if err != nil {
		log.Fatal(err)
	}
}
