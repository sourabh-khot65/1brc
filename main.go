package main

import (
	"fmt"
	"log"
	"time"

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

	// Simulate workload
	time.Sleep(5 * time.Second)

	fmt.Println("Application finished. CPU profile saved to benchmarks/cpu.prof")
}
