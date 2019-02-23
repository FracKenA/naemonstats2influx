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
var tags string

func init() {
	flag.StringVar(&filePath, "f", "", "Sets the input file.")
	flag.StringVar(&filePath, "file", "", "Sets the input file.")

	flag.StringVar(&measureName, "m", "naemon", "Sets the measurement name.")
	flag.StringVar(&measureName, "measurement-name", "naemon", "Sets the measurement name. (long option)")

	flag.StringVar(&tags, "t", "", "Sets any tags.")
	flag.StringVar(&tags, "tags", "", "Sets any tags. (long option)")
}

func main() {
	flag.Parse()

	logErr := log.New(os.Stderr, "", 0)
	logOut := log.New(os.Stdout, "", 0)

	if filePath == "" {
		logErr.Println("No file specified!")
		os.Exit(2)
	}

	logOut.Printf("Tags: %s\n", tags)

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
	logOut.Println(headers)
	logOut.Println(metrics)

	if err := fileScan.Err(); err != nil {
		logErr.Fatal(err)
	}

	if len(headers) != len(metrics) {
		logOut.Printf(
			"Header number (%d) does not match metric number (%d).",
			len(headers),
			len(metrics),
		)
	} else {
		logOut.Printf("%d headers, %d metrics", len(headers), len(metrics))
	}
}
