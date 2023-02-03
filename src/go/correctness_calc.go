//Lukas Diehm 1/27/22
package main

import (
	"os"
	"fmt"
	"log"
	"regexp"
	"strconv"
)



func main(){
	
	regex, _ := regexp.Compile("\"total_count\": [0-9]+") //Regex for finding count of issues in input file
	num_regex, _ := regexp.Compile("[0-9]+") //Regex for parsing count into only integer
	
	//closed issues
	data_closed, err1 := os.ReadFile("closed.txt")
	if(err1 != nil){
		fmt.Println("Did not find closed issues file from api")
		log.Fatal(err1)
	}
	closed_count := regex.FindString(string(data_closed))
	closed_count = num_regex.FindString(closed_count)

	//open issues
	data_open, err := os.ReadFile("open.txt")
	if(err != nil){
		fmt.Println("Did not find open issues file from api")
		log.Fatal(err)
	}
	open_count := regex.FindString(string(data_open))
	open_count = num_regex.FindString(open_count)


	score := calc_score(open_count, closed_count)
	fmt.Println(score)
	}

func calc_score(s1 string, s2 string) float64 {

	f1, err := strconv.ParseFloat(s1, 32)
	if err != nil {
		fmt.Println("Conversion of s1 to string float didn't work.")
	}
	f2, err1 := strconv.ParseFloat(s2, 32)
	if err1 != nil {
		fmt.Println("Conversion of s2 to string float didn't work.")
	}
	
	f3 := f2 / (f1 + f2)
	
	return f3
}