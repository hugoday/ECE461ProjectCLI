package main

import (
	// "fmt"
	"os"
	"testing"
	"io/ioutil"
)


func TestTestFunction(t *testing.T) {

	// name := "Gladys"
    // want := regexp.MustCompile(`\b`+name+`\b`)
    TestFunction()
    // if err != nil {
    //     t.Fatalf(`got err %v`, err)
    // }

}

// * START OF RESPONSIVENESS * \\

// NEED TO IMPLEMENT GITHUB TOKEN CALLING
// Function to get responsiveness metric score
func TestGetResponsiveness(t *testing.T) {

}

func TestRemoveScores(t *testing.T) {

}

// * END OF RESPONSIVENESS * \\

// * START OF RAMP-UP TIME * \\

// Function to get ramp-up time metric scor
func TestGetRampUpTime(t *testing.T) {

	
}

// * END OF RAMP-UP TIME * \\

// * START OF BUS FACTOR * \\

// Function to get bus factor metric score
func TestGetBusFactor(t *testing.T) {

	

}

// * END OF BUS FACTOR * \\

// * START OF CORRECTNESS * \\

// Function to get correctness metric score
func TestGetCorrectness(t *testing.T) {
	
}

func TestRunRestApi(t *testing.T) {

}

func TestTeardownRestApi(t *testing.T) {
	
}

func TestCalc_score(t *testing.T) {

}

// * END OF CORRECTNESS * \\

// * START OF LICENSE COMPATABILITY * \\

// Function to get license compatibility metric score
func TestGetLicenseCompatibility(t *testing.T) {
	
}

func TestSearchForLicenses(t *testing.T) {
	
}

func TestCheckFileForLicense(t *testing.T) {
	
}

// * END OF LICENSE COMPATABILITY * \\

// * START OF REPO CLONING/REMOVING  * \\

func TestCloneRepo(t *testing.T) {
	
}

func TestClearRepoFolder(t *testing.T) {
	
}

// * END OF REPO CLONING/REMOVING  * \\

// * START OF STDOUT * \\

func TestPrintRepo(t *testing.T) {
	
}

func TestRepoOUT(t *testing.T) {
	rescueStdout := os.Stdout
    read, w, _ := os.Pipe()
    os.Stdout = w

	var r repo
	r.URL = "testUrl"
	r.netScore = 7.4
	repoOUT(&r)

    w.Close()
    out, _ := ioutil.ReadAll(read)
    os.Stdout = rescueStdout
	// t.Errorf(string(out))
	// t.Errorf("hi")
	// t.Fail()

    if string(out) != "this is value: test" {
        t.Errorf("Expected %s, got %s", "this is value: test", out)
    }

	// fmt.Print("{\"URL\":\"", r.URL, "\", \"NET_SCORE\":", r.netScore, ", \"RAMP_UP_SCORE\":", r.rampUpTime, ",\"CORRECTNESS_SCORE\":", r.correctness, ", \"BUS_FACTOR_SCORE\":", r.busFactor, ", \"RESPONSIVE_MAINTAINER_SCORE\":", r.responsiveness, ", \"LICENSE_SCORE\":", r.licenseCompatibility, "}\n")
}

// * END OF STDOUT * \\

// * START OF SORTING * \\

func TestAddRepo(t *testing.T) {
	
}
