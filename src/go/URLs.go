package main

import (
	"bufio"
	"fmt"
	"math"

	"github.com/estebangarcia21/subprocess"

	// "github.com/go-git/go-git/v5"
	"log"
	"os"
	"os/exec"

	// "subprocess"
	// "fmt"
	// "log"
	// "os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	// "time"
)

// Struct for a repository
// Includes each metric, and the total score at the end
// this repo struct will be the input to the linked lists,
// where we will pass urls by accessing the repo's url
type repo struct {
	URL                  string
	responsiveness       int
	correctness          float64
	rampUpTime           int
	busFactor            int
	licenseCompatibility int
	totalScore           int
	next                 *repo
}

func runModule(function string) string {
	setup := "import sys; sys.path.append('../'); from src.python import test;"
	cmd := exec.Command("python", "-c", setup+"print("+function+")")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func runRestApi(url string) {

	index := strings.Index(url, ".com/")
	if index == -1 {
		fmt.Println("No '.com/' found in the string")
		return
	}
	url = url[index+5:]

	token := os.Getenv("GITHUB_TOKEN")
	code := `import os;
os.remove('src/python/issues/closed.txt') if os.path.exists('src/python/issues/closed.txt') else "continue";
os.remove('src/python/issues/open.txt') if os.path.exists('src/python/issues/open.txt') else "continue";
os.system('curl -i -H "Authorization: token ` + token + `" https://api.github.com/search/issues?q=repo:` + url + `+type:issue+state:closed >> src/python/issues/closed.txt');
os.system('curl -i -H "Authorization: token ` + token + `" https://api.github.com/search/issues?q=repo:` + url + `+type:issue+state:open >> src/python/issues/open.txt');`

	cmd := exec.Command("python", "-c", code)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	// setup := "\"import sys; sys.path.append('../'); from src.python import rest_api;"
	// // cmd := exec.Command("python", "-c", setup+"rest_api.getIssues(\\\""+url+"\\\")\"")

	// s := subprocess.New("python -c " + setup + "rest_api.getIssues(\\\"" + url + "\")\"")
	// fmt.Println(s)
	// if err := s.Exec(); err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println(err)
	// 	// eturn r("ERROR")
	// }
	// return string("SUCCESS")
	return
}

func teardownRestApi() {
	setup := "import sys; sys.path.append('../'); from src.python import rest_api;"
	cmd := exec.Command("python", "-c", setup+"rest_api.deleteIssues()")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

// Function to get responsiveness metric score
func getResponsiveness(url string) int {

	return -1
}

// Function to get correctness metric score
func getCorrectness(url string) float64 {
	fmt.Println("Getting correctness...")

	runRestApi(url)

	regex, _ := regexp.Compile("\"total_count\": [0-9]+") //Regex for finding count of issues in input file
	num_regex, _ := regexp.Compile("[0-9]+")              //Regex for parsing count into only integer

	//closed issues
	data_closed, err1 := os.ReadFile("./src/python/issues/closed.txt")
	if err1 != nil {
		fmt.Println("Did not find closed issues file from api, invalid url: " + url)
		log.Fatal(err1)
		return 0
	}
	closed_count := regex.FindString(string(data_closed))
	closed_count = num_regex.FindString(closed_count)

	//open issues
	data_open, err := os.ReadFile("./src/python/issues/open.txt")
	if err != nil {
		fmt.Println("Did not find open issues file from api, invalid url: " + url)
		log.Fatal(err)
		return 0
	}
	open_count := regex.FindString(string(data_open))
	open_count = num_regex.FindString(open_count)
	fmt.Println("Open: " + open_count + "\nClosed: " + closed_count)

	score := calc_score(open_count, closed_count)
	if math.IsNaN(score) {
		score = 0
	}
	// fmt.Println(score)

	teardownRestApi()
	fmt.Println("[CORRECTNESS DONE] ", score)
	fmt.Println()
	return score
}

// Function to get ramp-up time metric scor
func getRampUpTime(url string) int {

	return -1

}

// Function to get bus factor metric score
func getBusFactor(url string) int {

	return -1

}

// Function to get license compatibility metric score
func getLicenseCompatibility(url string) int {
	fmt.Println("Checking for license... ")

	cloneRepo(url)
	foundLicense := searchForLicenses("./src/repos/rnd/")
	clearRepoFolder()

	if foundLicense {
		fmt.Println("[LICENSE FOUND]")
		return 1
	}
	fmt.Println("[LICENSE NOT FOUND]")
	return 0
}

func checkFileForLicense(path string) bool {
	license := "LGPL-2.1"
	file, err := os.Open(path)
	if err != nil {
		fmt.Print("Coudln't open path ")
		fmt.Println(err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.Index(line, license); idx != -1 {
			fmt.Println("Found license in file:", path)
			return true
		}
	}
	return false
}

func searchForLicenses(folder string) bool {
	found := false
	//walk the repo looking for the license
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if found {
			return nil
		}
		if info.IsDir() { //skip .git, etc
			if info.Name()[0] == '.' {
				return filepath.SkipDir
			}
		} else {
			// fmt.Println("Searching for license in: " + path)
			found = checkFileForLicense(path)
		}
		return nil
	})

	//catch errors
	if err != nil {
		fmt.Println(err)
	}
	return found
}

func calc_score(s1 string, s2 string) float64 {

	f1, err := strconv.ParseFloat(s1, 32)
	// fmt.Println(s1)
	if err != nil {
		fmt.Println("Conversion of s1 to string float didn't work.")
	}
	// fmt.Println(s2)
	f2, err1 := strconv.ParseFloat(s2, 32)
	if err1 != nil {
		fmt.Println("Conversion of s2 to string float didn't work.")
	}

	f3 := f2 / (f1 + f2)

	return f3
}

// this is a function to utilize createing a new repo and initializing each metric within
func newRepo(url string) *repo {
	r := repo{URL: url}
	r.busFactor = getBusFactor(r.URL)
	r.correctness = getCorrectness(r.URL)
	r.licenseCompatibility = getLicenseCompatibility(r.URL)
	r.rampUpTime = getRampUpTime(r.URL)
	r.responsiveness = getResponsiveness(r.URL)
	r.totalScore = r.busFactor + int(r.correctness*20) + r.licenseCompatibility + r.rampUpTime + r.responsiveness
	return &r
}

func printRepo(r repo) {
	fmt.Println("URL: " + r.URL)
	fmt.Println("busFactor: ", r.busFactor)
	fmt.Println("correctness: ", r.correctness)
	fmt.Println("licenseCompatibility: ", r.licenseCompatibility)
	fmt.Println("rampUpTime: ", r.rampUpTime)
	fmt.Println("responsiveness: ", r.responsiveness)
	fmt.Println("totalScore: ", r.totalScore)
	fmt.Println()
	fmt.Println()
}

func cloneRepo(url string) string {
	s := subprocess.New("git clone " + url + " src/repos/rnd")
	if err := s.Exec(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return ("ERROR")
	}
	return string("SUCCESS")
}

func traverseList(next repo) {
	for {
		printRepo(next)
		if next.next == nil {
			break
		}
		next = *next.next
	}
}

func clearRepoFolder() {
	s := subprocess.New("rm -Force -r ", subprocess.Arg("src/repos/*"))
	fmt.Println(s.Exec())
}

func main() {
	clearRepoFolder()

	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	head := repo{}
	head.URL = "HEAD"
	prev := &head

	for scanner.Scan() {
		new := newRepo(scanner.Text())
		prev.next = new
		prev = new
	}

	traverseList(head)

	fmt.Println("[DONE]")

}
