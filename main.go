package main

import (
	"flag"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	_ "image/png" // Enables PNG decoding
	"os"
	"strings"
	"time"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to the out channel
		for _, p := range paths {
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				// Added error handling
				fmt.Fprintf(os.Stderr, "Error reading image %s: %v\n", p, err)
				continue // Skip this image if there's an error
			}
			job.Image = img
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			err := imageprocessing.WriteImage(job.OutPath, job.Image)
			if err != nil {
				// Added error handling
				fmt.Fprintf(os.Stderr, "Error writing image %s: %v\n", job.OutPath, err)
				out <- false
			} else {
				out <- true
			}
		}
		close(out)
	}()
	return out
}

func main() {
	// Added flag to control use of goroutines
	useGoroutines := flag.Bool("goroutines", true, "Process images using goroutines")
	flag.Parse()

	// Replaced the input image files with files of our choosing
	imagePaths := []string{
		"images/image1.png",
		"images/image2.png",
		"images/image3.png",
		"images/image4.png",
	}

	// Added code to measure processing time
	startTime := time.Now()

	if *useGoroutines {
		// Use pipeline with goroutines
		fmt.Println("Processing with goroutines")
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}
		}
	} else {
		// Process synchronously without goroutines
		fmt.Println("Processing without goroutines")
		for _, p := range imagePaths {
			outPath := strings.Replace(p, "images/", "images/output/", 1)
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading image %s: %v\n", p, err)
				continue
			}
			img = imageprocessing.Resize(img)
			img = imageprocessing.Grayscale(img)
			err = imageprocessing.WriteImage(outPath, img)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing image %s: %v\n", outPath, err)
				continue
			}
			fmt.Println("Success!")
		}
	}

	// Output total processing time
	fmt.Printf("Processing took %v\n", time.Since(startTime))
}
