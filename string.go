package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func timeToString(currentTime time.Time) string {
	return currentTime.String()
}

func osfileToSting(currentOsFile *os.File) string {
	file, err := os.Open(currentOsFile.Name())

	defer file.Close()

	fileinfo, err := file.Stat()

	if err != nil {
		fmt.Printf("Heyo, there's no file here!\n")
	}

	name := fileinfo.Name()
	return name
}

func slashSeperator(unslashed string) string {
	s := unslashed
	_, file := filepath.Split(s)
	return file
}
