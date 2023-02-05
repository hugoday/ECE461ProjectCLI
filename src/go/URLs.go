package main

import (
	"bufio"
	"fmt"
	"github.com/estebangarcia21/subprocess"
	// "github.com/go-git/go-git/v5"
	"log"
	"os"
	"os/exec"
	// "subprocess"
	// "fmt"
	// "log"
	// "os"
	"regexp"
	"strconv"
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
	// fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func runRestApi(url string) {
	setup := "import sys; sys.path.append('../'); from src.python import rest_api;"
	cmd := exec.Command("python", "-c", setup+"rest_api.getIssues("+url+")")
	// fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}

func teardownRestApi() {
	setup := "import sys; sys.path.append('../'); from src.python import rest_api;"
	cmd := exec.Command("python", "-c", setup+"rest_api.deleteIssues()")
	// fmt.Println(cmd.Args)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

// Function to get responsiveness metric score
func getResponsiveness(url string) int {

	return 5
}

// Function to get correctness metric score
func getCorrectness(url string) float64 {

	fmt.Println(url)
	runRestApi(url)

	regex, _ := regexp.Compile("\"total_count\": [0-9]+") //Regex for finding count of issues in input file
	num_regex, _ := regexp.Compile("[0-9]+")              //Regex for parsing count into only integer

	//closed issues
	fmt.Println(os.Getwd())
	data_closed, err1 := os.ReadFile("./src/python/issues/closed.txt")
	if err1 != nil {
		fmt.Println("Did not find closed issues file from api")
		log.Fatal(err1)
	}
	closed_count := regex.FindString(string(data_closed))
	closed_count = num_regex.FindString(closed_count)

	//open issues
	data_open, err := os.ReadFile("./src/python/issues/open.txt")
	if err != nil {
		fmt.Println("Did not find open issues file from api")
		log.Fatal(err)
	}
	open_count := regex.FindString(string(data_open))
	open_count = num_regex.FindString(open_count)

	score := calc_score(open_count, closed_count)
	// fmt.Println(score)

	teardownRestApi()
	return score
}

// Function to get ramp-up time metric scor
func getRampUpTime(url string) int {

	return 3

}

// Function to get bus factor metric score
func getBusFactor(url string) int {

	return 2

}

// Function to get license compatibility metric score
func getLicenseCompatibility(url string) int {

	return 1
}

func calc_score(s1 string, s2 string) float64 {

	f1, err := strconv.ParseFloat(s1, 32)
	fmt.Println(s1)
	if err != nil {
		fmt.Println("Conversion of s1 to string float didn't work.")
	}
	fmt.Println(s2)
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
		fmt.Println(next.URL)
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

	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)

	head := repo{}
	head.URL = "HEAD"
	prev := &head

	for scanner.Scan() {
		new := repo{}
		new.URL = scanner.Text()
		prev.next = &new
		prev = &new
	}

	// s := subprocess.New("../../../test1.sh")
	// fmt.Println(s.Exec())

	traverseList(head)

	url := head.next
	for url != nil {
		fmt.Println(newRepo(url.URL))
		url = url.next

		// clearRepoFolder()
	}

	// fmt.Println()
	// fmt.Println("Cloning...")
	// clearRepoFolder()

	// url := head.next
	// for url != nil {
	// 	fmt.Println(cloneRepo(url.URL))
	// 	url = url.next

	// 	clearRepoFolder()
	// }

	fmt.Println("[DONE]")

}
