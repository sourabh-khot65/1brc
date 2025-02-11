package profiling

import (
	"os"
	"runtime/pprof"
)

// StartCPUProfiling creates the benchmarks directory (if needed), opens the cpu profile file,
// and starts CPU profiling. It returns the file handle and any error encountered.
func StartCPUProfiling() (*os.File, error) {
	// Ensure the benchmarks directory exists.
	if err := os.MkdirAll("benchmarks", 0755); err != nil {
		return nil, err
	}

	// Create or truncate the cpu profile file.
	f, err := os.Create("benchmarks/cpu.prof")
	if err != nil {
		return nil, err
	}

	// Start CPU profiling; if an error occurs, close the file.
	if err := pprof.StartCPUProfile(f); err != nil {
		f.Close()
		return nil, err
	}

	return f, nil
}

// StopCPUProfiling stops the CPU profiling.
func StopCPUProfiling() {
	pprof.StopCPUProfile()
}
