# 1brc
Exploring efficient data processing techniques to handle 1 billion rows using Golang, concurrency, and performance optimizations

Welcome to **1brc**, a Golang project that serves as a starting point for data processing and analysis challenges. This repository is structured to provide clear separation of concerns and facilitate both development and testing.

## Project Structure

- **benchmarks/**: Contains benchmark results and profiling reports.
- **data/**: Holds sample datasets and data generation scripts.
- **docs/**: Includes technical write-ups, optimizations, and findings.
- **internal/**: Core logic that is modular and reusable.
  - **processing/**: Data processing logic such as parsing and aggregations.
  - **storage/**: Storage and indexing utilities.
- **scripts/**: Helper scripts for automation, data loading, and profiling.
- **tests/**: Unit and performance tests.
- **main.go**: Entry point for running the challenge.
- **go.mod**: Go module file.
- **.gitignore**: Specifies files and directories to be ignored by Git.

## Getting Started

To get started with 1brc:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/sourabh-khot65/1brc.git
   cd 1brc
