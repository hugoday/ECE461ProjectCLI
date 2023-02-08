package main

import (
	"os"
	"fmt"
	"log"
	"regexp"
	"strings"
	"strconv"	
)
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
	totalScore           float64
}

// this is a function to utilize createing a new repo and initializing each metric within
func newRepo(url string) *repo {
	r := repo{URL: url}
	r.busFactor = getBusFactor(r.URL)
	r.correctness = getCorrectness(r.URL)
	r.licenseCompatibility = getLicenseCompatibility(r.URL)
	r.rampUpTime = getRampUpTime(r.URL)
	r.responsiveness = getResponsiveness(r.URL)
	r.totalScore = r.busFactor + r.correctness + r.licenseCompatibility + r.rampUpTime + r.responsiveness
	return &r
}

// Function to get responsiveness metric score
func getResponsiveness(url string) int {

	return 5
}

// Function to get correctness metric score
func getCorrectness(url string) int {

	return 4
}

// Function to get ramp-up time metric scor
func getRampUpTime(url string) int {

	return 3

}

// Function to get bus factor metric score
func getBusFactor(url string) float64 {

	regex, _ := regexp.Compile("[0-9]+") //Regex for parsing count into only integer

	short_log_raw_data, err1 := os.ReadFile("shortlog.txt")
	if err1 != nil {
		fmt.Println("Did not find closed issues file from api")
		log.Fatal(err1)
	}
	arr := strings.Split(string(short_log_raw_data), "\n")
	len_log := len(arr) - 1
	var num_bus_committers int
	if len_log < 100{
		num_bus_committers = 1
	}else{
		num_bus_committers = len_log / 100
	}
	total := 0
	total_bus_guys := 0
	var num string
	for i := 0; i < len_log; i++ {
		
		num = regex.FindString(arr[i])
		num_int, err2 := strconv.Atoi(num)
		if err2 != nil{
			fmt.Println("Conversion from string to int didn't work (bus factor calc)")
			log.Fatal(err2)
		}
		total += num_int
		if i >= len_log - num_bus_committers{
			total_bus_guys += num_int
		}
	}

	metric := (float64(total) - float64(total_bus_guys)) / float64(total)

	fmt.Println(metric) //return statement
}

// Function to get license compatibility metric score
func getLicenseCompatibility(url string) int {

	return 1
}

func main() {
	var r *repo
	r = newRepo("github.com")

	fmt.Println(r)
}
