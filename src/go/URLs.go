package main

// package imports

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	// "log"
	// "reflect"

	"github.com/estebangarcia21/subprocess"
)

// * START OF REPO STRUCTS * \\

// Struct for a repository
// Includes each metric, and the total score at the end
// this repo struct will be the input to the linked lists,
// where we will pass urls by accessing the repo's url
type repo struct {
	URL                  string
	responsiveness       float64
	correctness          float64
	rampUpTime           float64
	busFactor            float64
	licenseCompatibility float64
	netScore             float64
	next                 *repo
}

// this is a function to utilize createing a new repo and initializing each metric within
func newRepo(url string) *repo {
	InfoLogger.Println("Getting metrics for new repo ", url)
	cloneRepo(url)
	r := repo{URL: url}
	r.busFactor = getBusFactor(r.URL)
	r.correctness = getCorrectness(r.URL)
	r.licenseCompatibility = getLicenseCompatibility(r.URL)
	r.rampUpTime = getRampUpTime(r.URL)
	r.responsiveness = getResponsiveness(r.URL)
	if (r.busFactor == -1) || (r.correctness == -1) || (r.responsiveness == -1) || (r.rampUpTime == -1) || (r.licenseCompatibility == -1) {
		r.netScore = -1
	} else {
		r.netScore = ((75 * r.licenseCompatibility) + (15 * r.busFactor) + (20 * r.responsiveness) + (20 * r.rampUpTime) + (20 * r.correctness)) / 150
	}
	clearRepoFolder()

	InfoLogger.Println("Done getting metrics for ", url)

	return &r
}

// * END OF REPO STRUCTS * \\

// * START OF RESPONSIVENESS * \\

// Function to get responsiveness metric score
func getResponsiveness(url string) float64 {
	InfoLogger.Println("Getting responsiveness for ", url)
	// Variable declarations
	var command string

	// command to run API.py with a URL passed in, output is stored in score.txt
	command = "python3 src/python/API.py \"" + url + "\" >> src/metric_scores/responsiveness/score.txt"

	// creates process for command on shell
	s := subprocess.New(command, subprocess.Shell)

	// executes command
	s.Exec()

	// opens score.txt
	file, _ := os.Open("src/metric_scores/responsiveness/score.txt")

	// create scanner to scan file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// reads string in score txt, converts to float64
		s := scanner.Text()
		score, err := strconv.ParseFloat(s, 64)
		// if error return 0 as score and remove file, otherwise return score
		if err != nil {
			removeScores()
			DebugLogger.Println("Error pasing score : ", err)
			return 0
		} else {
			removeScores()
			InfoLogger.Printf("Got responsiveness score: %f for %s\n", score, url)
			return score
		}
	}
	// in case of error, removeScores, return 0
	removeScores()
	DebugLogger.Println("Error getting responsiveness")
	return 0
}

// This function removes the responsiveness score text file
func removeScores() {
	DebugLogger.Println("Removing scores")
	
	// Variable declarations
	var command string

	// command to run API.py with a URL passed in, output is stored in score.txt
	command = "python3 -c 'import os; os.remove(\"src/metric_scores/responsiveness/score.txt\") if os.path.exists(\"src/metric_scores/responsiveness/score.txt\") else \"continue\";'"

	// creates process for command on shell
	r := subprocess.New(command, subprocess.Shell)

	// executes command
	r.Exec()
	DebugLogger.Println("Scores removed")
}

// * END OF RESPONSIVENESS * \\

// * START OF RAMP-UP TIME * \\

// Function to get ramp-up time metric score, calls rampUpTime.py, reads result from RU_Result.txt, returns that result as float
func getRampUpTime(url string) float64 {
	InfoLogger.Println("Getting ramp up time for ", url)
	var command string
	command = "python3 src/python/rampUpTime.py"
	r := subprocess.New(command, subprocess.Shell)
	r.Exec()
	dat, err := os.ReadFile("src/metric_scores/rampuptime/RU_Result.txt")
	if err != nil {
		DebugLogger.Println("File open failed")
	}
	command = "rm src/metric_scores/rampuptime/RU_Result.txt"
	r = subprocess.New(command, subprocess.Shell)
	r.Exec()
	f1, err := strconv.ParseFloat(string(dat), 64)
	if err != nil {
		DebugLogger.Println("Conversion of string to float didn't work.")
	}
	InfoLogger.Printf("Got ramp up time score: %f for %s\n", f1, url)
	return f1

}

// * END OF RAMP-UP TIME * \\

// * START OF BUS FACTOR * \\

// Function to get bus factor metric score
func getBusFactor(url string) float64 {
	InfoLogger.Println("Getting bus factor for ", url)
	
	make_shortlog_file()
	regex, _ := regexp.Compile("[0-9]+") //Regex for parsing count into only integer

	short_log_raw_data, err1 := os.ReadFile("src/metric_scores/busfactor/shortlog.txt")
	if err1 != nil {
		DebugLogger.Println("Did not find shortlog file")
		log.Fatal(err1)
	}

	arr := strings.Split(string(short_log_raw_data), "\n") // parsing shortlog file by lines

	len_log := len(arr) - 1

	if len_log < 1 {
		DebugLogger.Println("No committers for repo")
		delete_shortlog_file()
		return 0
	}

	var num_bus_committers int
	if len_log < 100 {
		num_bus_committers = 1
	} else {
		num_bus_committers = len_log / 100
	}

	total := 0
	total_bus_guys := 0
	var num string


	for i := 0; i < len_log; i++ {
		num = regex.FindString(arr[i])
		num_int, err2 := strconv.Atoi(num)
		if err2 != nil {
			DebugLogger.Println("Conversion from string to int didn't work (bus factor calc)")
			log.Fatal(err2)
		}
		total += num_int
		if i < num_bus_committers {
			total_bus_guys += num_int
		}
	}
	delete_shortlog_file()
	metric := (float64(total) - float64(total_bus_guys)) / float64(total)

	InfoLogger.Printf("Got bus factor score: %f for %s\n", metric, url)
	return metric
}

func make_shortlog_file() {
	DebugLogger.Println("Making shortlog file")
	os.Chdir("src/metric_scores/repos")

	cmd := exec.Command("git", "shortlog", "HEAD", "-se", "-n")

	out, err := cmd.Output()

	if err != nil {
		DebugLogger.Println("Did not find closed issues file from api, invalid url: ")
	}
	os.Chdir("../")
	os.Chdir("busfactor")
	os.WriteFile("shortlog.txt", out, 0644)
	os.Chdir("../")
	os.Chdir("../")
	os.Chdir("../")

	DebugLogger.Println("Done making shortlog file")

}

func delete_shortlog_file() {
	DebugLogger.Println("Deleting shortlog file")
	var command string
	command = "rm -f src/metric_scores/busfactor/shortlog.txt"
	s := subprocess.New(command, subprocess.Shell)
	s.Exec()

	DebugLogger.Println("Done deleting shortlog file")
}

// * END OF BUS FACTOR * \\

// * START OF CORRECTNESS * \\

// Function to get correctness metric score
func getCorrectness(url string) float64 {
	InfoLogger.Println("Getting correctness for ", url)

	// runs RestAPI on url
	runRestApi(url)

	regex, _ := regexp.Compile("\"total_count\": [0-9]+") //Regex for finding count of issues in input file
	num_regex, _ := regexp.Compile("[0-9]+")              //Regex for parsing count into only integer

	//gets # of closed issues
	data_closed, err1 := os.ReadFile("./src/metric_scores/correctness/closed.txt")
	closed_count := regex.FindString(string(data_closed))
	closed_count = num_regex.FindString(closed_count)

	//gets # of open issues
	data_open, err := os.ReadFile("./src/metric_scores/correctness/open.txt")
	open_count := regex.FindString(string(data_open))
	open_count = num_regex.FindString(open_count)
	if err != nil || err1 != nil {
		fmt.Println("Did not find issues file from api, invalid url: " + url)
		return 0
	}

	// calculates correctness score
	score := calc_score(open_count, closed_count)
	if math.IsNaN(score) {
		score = 0
	}

	// removes files API output files
	teardownRestApi()

	InfoLogger.Printf("Got correctness: %f for %s\n", score, url)
	// returns score
	return score
}

// Runs rest api
func runRestApi(url string) int {
	DebugLogger.Println("Running rest API on ", url)

	index := strings.Index(url, ".com/")
	if index == -1 {
		// fmt.Println("No '.com/' found in the string")
		return 1
	}
	url = url[index+5:]

	// gets env github token
	token := os.Getenv("GITHUB_TOKEN")
	DebugLogger.Println("Got GIHUB_TOKEN ", token)

	// command to check if files are in directory, removes them, then runs restapi
	command := "python3 -c 'import os; os.remove(\"src/metric_scores/correctness/closed.txt\") if os.path.exists(\"src/metric_scores/correctness/closed.txt\") else \"continue\"; os.remove(\"src/metric_scores/correctness/open.txt\") if os.path.exists(\"src/metric_scores/correctness/open.txt\") else \"continue\"; os.system(\"curl -s -i -H \\\"Authorization: token " + token + "\\\" https://api.github.com/search/issues?q=repo:" + url + "+type:issue+state:closed >> src/metric_scores/correctness/closed.txt\"); os.system(\"curl -s -i -H \\\"Authorization: token " + token + "\\\" https://api.github.com/search/issues?q=repo:" + url + "+type:issue+state:open >> src/metric_scores/correctness/open.txt\");'"

	// creates subprocess to run on shell
	r := subprocess.New(command, subprocess.Shell)

	//Executes subprocess
	r.Exec()

	DebugLogger.Println("Done running rest API on ", url)
	return 0
}

func teardownRestApi() {
	DebugLogger.Println("Tearing down rest API")
	// Variable declarations
	var command string

	// removes correctness open and closed issue files.
	command = "python3 -c 'import os; os.remove(\"src/metric_scores/correctness/closed.txt\") if os.path.exists(\"src/metric_scores/correctness/closed.txt\") else \"continue\"; os.remove(\"src/metric_scores/correctness/open.txt\") if os.path.exists(\"src/metric_scores/correctness/open.txt\") else \"continue\";'"

	// creates process for command on shell
	r := subprocess.New(command, subprocess.Shell)

	// executes command
	r.Exec()
	
	DebugLogger.Println("Done tearing down rest API")
}

// calculates score for correctness
func calc_score(s1 string, s2 string) float64 {
	DebugLogger.Println("Converting string to float")

	// converts string to float
	f1, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		DebugLogger.Println("Conversion of s1 to string float didn't work.")
	}
	//converts string to float
	f2, err1 := strconv.ParseFloat(s2, 32)
	if err1 != nil {
		DebugLogger.Println("Conversion of s2 to string float didn't work.")
	}

	// round(20 * closed issues / (open + closed issues)) / 20 = correctness score [0,1]
	f3 := f2 / (f1 + f2)

	DebugLogger.Println("Done onverting string to float")
	return f3
}

// * END OF CORRECTNESS * \\

// * START OF LICENSE COMPATABILITY * \\

// Function to get license compatibility metric score
func getLicenseCompatibility(url string) float64 {
	InfoLogger.Println("Getting license compatibility for ", url)

	// checks to see if license is found
	foundLicense := searchForLicenses("./src/metric_scores/repos/")
	

	InfoLogger.Printf("Got license compatibility: %f for %s\n", foundLicense, url)
	// returns score based on if license is found
	return foundLicense
}

// searches directories for license
func searchForLicenses(folder string) float64 {
	DebugLogger.Println("Searching for licences in ", folder)
	// initialize found as false
	found := false

	//walk the repo looking for the license
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if found {
			return nil
		}
		if info == nil {
			return nil
		}
		if info.IsDir() {
			if len(info.Name()) > 0 && info.Name()[0] == '.' {
				return filepath.SkipDir
			}
		} else {
			// checks if file has license
			found = checkFileForLicense(path)
		}
		return nil
	})

	//catch errors
	if err != nil {
		DebugLogger.Println("Error searching for licenses: ", err)
		return 0.00
	}

	if found {
		DebugLogger.Println("Done searching for licence: found")
		return 1.00
	}
	DebugLogger.Println("Done searching for licence: not found")
	return 0.00
}

// searches file for license
func checkFileForLicense(path string) bool {
	DebugLogger.Println("Checking file for license: ", path)
	// license to search for
	license := "LGPL-2.1"

	// searches for license string
	file, err := os.Open(path)
	if err != nil {
		DebugLogger.Println("Error opening file: ", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.Index(line, license); idx != -1 {
			DebugLogger.Println("Found license in file:", path)
			return true
		}
	}
	DebugLogger.Println("No license found in file: ", path)
	return false
}

// * END OF LICENSE COMPATABILITY * \\

// * START OF REPO CLONING/REMOVING  * \\

func cloneRepo(url string) string {
	DebugLogger.Println("Cloning repo: ", url)
	s := subprocess.New("git clone --quiet "+url+" src/metric_scores/repos", subprocess.Shell)
	if err := s.Exec(); err != nil {
		log.Fatal(err)
		DebugLogger.Printf("Error cloning repo %s: %s\n", url, err)
		return ("ERROR CLONING")
	}
	index := strings.Index(url, ".com/")
	if index == -1 {
		DebugLogger.Println("No '.com/' found in the string ", url)
		return "FAILURE"
	}

	r, _ := regexp.Compile("/")
	a := r.Split(url[index+5:], 2)

	DebugLogger.Println("Cloned repo: ", a[1])
	return a[1]
}

func clearRepoFolder() {
	DebugLogger.Println("Clearing repo folder")
	s := subprocess.New("rm -rf ", subprocess.Arg("src/metric_scores/repos"), subprocess.Shell)
	s.Exec()
	DebugLogger.Println("Done clearing repo folder")
}

// * END OF REPO CLONING/REMOVING  * \\

// * START OF STDOUT * \\

func printRepo(next *repo) {
	for {
		if next.URL != "temp" {
			repoOUT(next)
		}
		if next.next == nil {
			break
		}
		next = next.next
	}
}

// Prints each repo in NDJSON output format (.2 decimals for floats)
func repoOUT(r *repo) {
	InfoLogger.Println("Final repo result:")
	InfoLogger.Printf("{\"URL\":\"%s\", \"NET_SCORE\":%.2f, \"RAMP_UP_SCORE\":%.2f, \"CORRECTNESS_SCORE\":%.2f, \"BUS_FACTOR_SCORE\":%.2f, \"RESPONSIVE_MAINTAINER_SCORE\":%.2f, \"LICENSE_SCORE\":%.2f} \n", r.URL, r.netScore, r.rampUpTime, r.correctness, r.busFactor, r.responsiveness, r.licenseCompatibility)
	fmt.Printf("{\"URL\":\"%s\", \"NET_SCORE\":%.2f, \"RAMP_UP_SCORE\":%.2f, \"CORRECTNESS_SCORE\":%.2f, \"BUS_FACTOR_SCORE\":%.2f, \"RESPONSIVE_MAINTAINER_SCORE\":%.2f, \"LICENSE_SCORE\":%.2f} \n", r.URL, r.netScore, r.rampUpTime, r.correctness, r.busFactor, r.responsiveness, r.licenseCompatibility)
}

// * END OF STDOUT * \\

// * START OF SORTING * \\

func addRepo(head *repo, curr *repo, temp *repo) *repo {
	DebugLogger.Println("Adding repo to list")
	head.next = curr
	if curr == nil {
		head.next = temp
	} else {
		if curr.netScore >= temp.netScore {
			curr = addRepo(curr, curr.next, temp)
		} else {
			head.next = temp
			temp.next = curr
		}
	}

	DebugLogger.Println("Done adding repo to list")
	return head
}

// * END OF SORTING * \\
