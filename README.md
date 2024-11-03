# Image Processing Pipeline with / without Goroutines

This repository contains an image processing pipeline implemented in Go, demonstrating the use of goroutines and channels for concurrent processing. The pipeline reads images, resizes them, converts them to grayscale, and saves the processed images to an output directory. The program can be run with or without goroutines to compare processing times.

## Introduction

This project showcases how Go's concurrency model can be leveraged to improve the performance of an image processing pipeline. By using goroutines and channels, the pipeline stages can process images concurrently, resulting in faster overall processing times compared to a synchronous approach.

## Test Results
Average time consumed using goroutines: 72.67ms
Average time consumed without using goroutines: 106.57ms

## Features

- **Concurrent Image Processing**: Utilize goroutines for each stage of the pipeline to process images concurrently.
- **Error Handling**: Added robust error checking for image file input and output operations.
- **Benchmarking**: Measure and compare processing times with and without goroutines.
- **Configurable Execution**: Run the program with or without goroutines using a command-line flag.
- **Unit Tests**: Includes unit tests for image processing functions to ensure correctness.

## Usage

### Running with Goroutines

To run the image processing pipeline using goroutines:

```bash
go run main.go -goroutines=true
```

### Running without Goroutines

To run the pipeline synchronously without goroutines:

```bash
go run main.go -goroutines=false
```

### Command-Line Flags

- `-goroutines` (default `true`): Specifies whether to use goroutines for processing.

## Unit Tests

Unit tests are included to verify the correctness of the image processing functions.

### Running Tests

```bash
go test ./...
```

### Test Coverage

Ensure all functions are properly tested, and consider using coverage tools to identify untested code paths.

## Code Modifications

Significant changes and enhancements made to the original code include:

1. **Error Handling for Image I/O**

   - Modified `ReadImage` and `WriteImage` functions to return errors instead of panicking.
   - Updated pipeline stages to handle errors gracefully, logging them and skipping problematic images.

2. **Custom Image Files**

   - Replaced the original input image files with custom images located in the `images/` directory.
   - Ensure that your images are placed in the `images/` directory before running the program.

3. **Command-Line Flag for Goroutines**

   - Added a `-goroutines` flag to toggle between concurrent and synchronous execution.
   - Allows easy comparison of processing times and understanding the impact of concurrency.

4. **Benchmarking Implementation**

   - Incorporated time measurement using Go's `time` package.
   - Outputs the total processing time after the pipeline completes.

5. **Additional Enhancements**

   - Improved logging for better clarity on successes and failures.
   - Ensured the pipeline continues processing remaining images even if some fail.
   - Organized code for better readability and maintenance.

## Project Structure

```
image-processing-pipeline/
├── main.go
├── image_processing/
│   ├── image_processing.go
│   └── image_processing_test.go
├── images/
│   ├── my_image1.jpg
│   ├── my_image2.jpg
│   ├── my_image3.jpg
│   └── my_image4.jpg
├── images/output/
│   ├── my_image1.jpg
│   ├── my_image2.jpg
│   ├── my_image3.jpg
│   └── my_image4.jpg
├── go.mod
└── README.md
```

- **main.go**: Entry point of the application containing the pipeline setup.
- **image_processing/**: Package containing image processing functions and tests.
- **images/**: Directory containing input images.
- **images/output/**: Directory where processed images are saved.
- **go.mod**: Go module file.