package main

import (
	"bufio"
	"os"
)

// Main
func main() {
	// Makes sure repository folder is clear
	clearRepoFolder()

	// Opens URL file and creates a scanner
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	// Create head and temporary repo nodes
	var head *repo
	var hold *repo
	head = &repo{URL: "HEAD"}

	// for each url in the file
	for scanner.Scan() {
		//Create new repositories with current URL scanned
		hold = newRepo(scanner.Text())
		// Adds repository to Linked List in sorted order by net score
		head = addRepo(head, head.next, hold)
	}

	// Prints each repository in NDJSON format to stdout (sorted highest to low based off net score)
	printRepo(head.next)
}
