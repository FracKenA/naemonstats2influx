package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var filePath string
var measureName string
var tagsExtra string

func init() {
	flag.StringVar(&filePath, "f", "", "Sets the input file.")
	flag.StringVar(&filePath, "file", "", "Sets the input file.")

	flag.StringVar(&measureName, "m", "naemon", "Sets the measurement name.")
	flag.StringVar(&measureName, "measurement-name", "naemon", "Sets the measurement name. (long option)")

	flag.StringVar(&tagsExtra, "t", "", "Sets any extra tags.")
	flag.StringVar(&tagsExtra, "tags-extra", "", "Sets any extra tags. (long option)")
}

func main() {
	flag.Parse()

	logErr := log.New(os.Stderr, "", 0)
	logOut := log.New(os.Stdout, "", 0)

	if filePath == "" {
		logErr.Println("No file specified!")
		os.Exit(2)
	}

	fileHandle, err := os.Open(filePath)
	if err != nil {
		logErr.Printf("Error %s", err)
		os.Exit(3)
	}
	defer fileHandle.Close()

	fileScan := bufio.NewScanner(fileHandle)
	// Ti's expected that the file will consist of two lines in a CSV format.
	// Line 1 is the headers and line 2 is the data.
	// Starting the to read the file.
	fileScan.Scan()
	headers := strings.Split(
		strings.Trim(fileScan.Text(), ", "),
		",",
	)
	// Advancing the marker to the second line.
	fileScan.Scan()
	metrics := strings.Split(
		strings.Trim(fileScan.Text(), ", "),
		",",
	)

	if err := fileScan.Err(); err != nil {
		logErr.Fatal(err)
	}

	for index, metricName := range headers {
		if tagsExtra != "" {
			logOut.Printf("%s,tags=%s,%s %s=%s\n", measureName, metricName, tagsExtra, metricName, metrics[index])
		} else {
			logOut.Printf("%s,tags=%s %s=%s\n", measureName, metricName, metricName, metrics[index])
		}
	}
}
