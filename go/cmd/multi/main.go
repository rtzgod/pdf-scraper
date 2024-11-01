package main

import (
	"fmt"
	internal2 "github.com/rtzgod/email-scraper/go/internal"
	"os"
	"regexp"
	"sync"
)

const (
	emailRegex = `[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`
)

func main() {
	fileNames, err := internal2.DirFiles("pdfs")
	fmt.Printf("Total files found: %d\n", len(fileNames))
	if err != nil {
		fmt.Println(err)
	}

	outputFileName := "output.txt"
	var wg sync.WaitGroup
	re := regexp.MustCompile(emailRegex)

	outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for _, file := range fileNames {
		wg.Add(1)

		go func(file os.FileInfo) {
			defer wg.Done()
			fmt.Println("Starting to process file:", file.Name())

			text, err := internal2.ReadPDF("pdfs/" + file.Name())
			if err != nil {
				return
			}

			emails := re.FindAllString(text, -1)

			for _, email := range emails {
				if _, err := outputFile.WriteString(email + "\n"); err != nil {
					fmt.Println("Failed to write to file:", err)
					return
				}
			}
		}(file)
	}

	wg.Wait()
}
