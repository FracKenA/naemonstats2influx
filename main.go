package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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
		logErr.Fatal("No file specified!")
		os.Exit(2)
	}

	logOut.Printf("Tags: %s\n", tags)

	fileData, err := os.Open(filePath)
	if err != nil {
		logErr.Fatalf("Error %s", err)
		os.Exit(3)
	}
	defer fileData.Close()

	fileScan := bufio.NewScanner(fileData)
	for fileScan.Scan() {
		logOut.Println(fileScan.Text())
	}

	if err := fileScan.Err(); err != nil {
		logErr.Fatal(err)
	}
}
