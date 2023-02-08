package main

import (
	"os"
	"fmt"
	"log"
	"regexp"
	"strings"
	"strconv"
)

func main(){

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

func print (in string){
	fmt.Println(in)
}