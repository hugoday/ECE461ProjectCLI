package main

import (
	"bufio"
	// "fmt"
	"log"
	// "math"
	// "utils"
	"fmt"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
)

type NoLog int

func (NoLog) Write([]byte) (int, error) {
	return 0, nil
}

func main() {

	fmt.Println("Aditya Srikanth")

	doLogging := true
	logFileName := os.Getenv("LOG_FILE")
	logLevel := os.Getenv("LOG_LEVEL")
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil || (logLevel != "1" && logLevel != "2") {
		doLogging = false
	}

	if doLogging {
		if logLevel == "2" {
			DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		} else {
			DebugLogger = log.New(new(NoLog), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		}
		InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		DebugLogger = log.New(new(NoLog), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
		InfoLogger = log.New(new(NoLog), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	// Makes sure repository folder is clear
	clearRepoFolder()

	// Opens URL file and creates a scanner
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	// Create head and temporary repo nodes
	var head *repo
	var hold *repo
	head = &repo{URL: "HEAD"}

	InfoLogger.Println("Beginning URL file read")
	// for each url in the file
	for scanner.Scan() {
		//Create new repositories with current URL scanned
		hold = newRepo(scanner.Text())
		InfoLogger.Println("New repo created successfully")
		// Adds repository to Linked List in sorted order by net score
		head = addRepo(head, head.next, hold)
	}

	// Prints each repository in NDJSON format to stdout (sorted highest to low based off net score)
	printRepo(head.next)
}
