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
   ```

2. **Run the application:**
   The application requires a filename as a command-line argument. You can run the application using the following command:
   ```sh
   go run main.go <filename>
   ```
   Replace `<filename>` with the path to the data file you want to process. The data file should contain city names and temperatures in the following format:
   ```
   CityName;Temperature
   ```

3. **Output:**
   After processing the data, the application will print the minimum, maximum, average temperatures, and count of entries for each city to the console.

### Example Data File
Your data file should look something like this:

```
New York;30
Los Angeles;25
New York;35
Chicago;20
Los Angeles;28
```

### Approach (Attempt 1)
In this initial attempt, the project focuses on efficiently processing temperature data for various cities. The approach includes:

- **Data Parsing**: The application reads a file containing city names and their corresponding temperatures, ensuring that the data is correctly formatted.
- **Temperature Calculations**: For each city, the application calculates the minimum, maximum, and average temperatures, as well as the count of temperature entries.
- **Error Handling**: The implementation includes error handling to manage malformed lines in the input data and parsing errors.
- **Output**: The results are printed to the console, providing a clear summary of the temperature statistics for each city.

This approach serves as a foundational step towards building a more robust data processing system capable of handling larger datasets and more complex analyses in future iterations.