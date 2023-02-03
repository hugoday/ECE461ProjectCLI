// package main

// import "fmt"

// // Struct for a repository
// // Includes each metric, and the total score at the end
// // this repo struct will be the input to the linked lists,
// // where we will pass urls by accessing the repo's url
// type repo struct {
// 	URL                  string
// 	responsiveness       int
// 	correctness          int
// 	rampUpTime           int
// 	busFactor            int
// 	licenseCompatibility int
// 	totalScore           int
// }

// // this is a function to utilize createing a new repo and initializing each metric within
// func newRepo(url string) *repo {
// 	r := repo{URL: url}
// 	r.busFactor = getBusFactor(r.URL)
// 	r.correctness = getCorrectness(r.URL)
// 	r.licenseCompatibility = getLicenseCompatibility(r.URL)
// 	r.rampUpTime = getRampUpTime(r.URL)
// 	r.responsiveness = getResponsiveness(r.URL)
// 	r.totalScore = r.busFactor + r.correctness + r.licenseCompatibility + r.rampUpTime + r.responsiveness
// 	return &r
// }

// // Function to get responsiveness metric score
// func getResponsiveness(url string) int {

// 	return 5
// }

// // Function to get correctness metric score
// func getCorrectness(url string) int {

// 	return 4
// }

// // Function to get ramp-up time metric scor
// func getRampUpTime(url string) int {

// 	return 3

// }

// // Function to get bus factor metric score
// func getBusFactor(url string) int {

// 	return 2

// }

// // Function to get license compatibility metric score
// func getLicenseCompatibility(url string) int {

// 	return 1
// }

// func main() {
// 	var r *repo
// 	r = newRepo("github.com")

// 	fmt.Println(r)
// }
