package main

import (
	"fmt"
	internal2 "github.com/rtzgod/email-scraper/go/internal"
	"os"
	"regexp"
)

const (
	emailRegex = `(?i)\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}\b`
)

func main() {
	filePaths, err := internal2.DirFiles("pdfs")
	fmt.Printf("Total files found: %d\n", len(filePaths))
	if err != nil {
		fmt.Println(err)
	}

	successfulReadingsCounter := 0
	outputFileName := "output.txt"
	re := regexp.MustCompile(emailRegex)

	outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for _, path := range filePaths {
		text, err := internal2.ReadPDF("pdfs/" + path.Name())
		if err != nil {
			if err.Error() == "unable to get plain text" {
				fmt.Println("Skipping file due to error:", err)
				continue
			}
			fmt.Println(err)
			return
		}

		successfulReadingsCounter++
		emails := re.FindAllString(text, -1)

		for _, email := range emails {
			if _, err := outputFile.WriteString(email + "\n"); err != nil {
				fmt.Println("Failed to write to file:", err)
				return
			}
		}
	}
	fmt.Printf("Total files read: %d\n", successfulReadingsCounter)
}
