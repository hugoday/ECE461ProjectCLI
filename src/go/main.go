package main

import (
	"bufio"
	// "fmt"
	"log"
	// "math"
	// "utils"
	"os"
	// "os/exec"
	// "path/filepath"
	// "regexp"
	// "strconv"
	// "strings"

	// "github.com/estebangarcia21/subprocess"
	// "subprocess"
	// "fmt"
	// "log"
	// "os"
	// "github.com/go-git/go-git/v5"
	// "time"
)

var (
    DebugLogger   *log.Logger
    InfoLogger    *log.Logger
)


func main() {
	logFileName := os.Getenv("LOG_FILE")
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    DebugLogger = log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Makes sure repository folder is clear
	clearRepoFolder()

	// Opens URL file and creates a scanner
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	// Create head and temporary repo nodes
	var head *repo
	var hold *repo
	head = &repo{URL: "HEAD"}

	for scanner.Scan() {
		//Create new repositories with current URL scanned
		hold = newRepo(scanner.Text())
		head = addRepo(head, head.next, hold)
		// Add New Repo to linked list
		// NEEDS TO BE REPLACED WITH SORTING METHOD
	}

	//
	printRepo(head.next)
}