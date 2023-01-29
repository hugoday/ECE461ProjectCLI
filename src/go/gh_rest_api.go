//Lukas Diehm 1/27/22
package main

import (
	"os"
	"fmt"
	"log"
	"github.com/estebangarcia21/subprocess"
)

func main(){

	issues_file_name := "issues.txt"
	make_issues_file(issues_file_name)
	
} 


func make_issues_file(text_file_name string){
	//Initializing relevant CLI commands
	str := "gh api --method GET /repos/octocat/Spoon-Knife/issues >> " + text_file_name
	fmt.Println(str)

	github_rest_api_CMD := subprocess.New("cat Repo.go")
	// github_rest_api_CMD := subprocess.New(str)
	str2 := "rm " + text_file_name
	delete_file_CMD := subprocess.New(str2)
	
	// deleting output file if it exists
	if _, err0 := os.Stat(text_file_name); err0 == nil {
		fmt.Println("deleting existing file")
		
		err1 := delete_file_CMD.Exec()
		if err1 != nil {
			fmt.Println("failed delete cmd")
			
			log.Fatal(err1)		
		}
	}
		
	fmt.Println(github_rest_api_CMD.Exec())
	fmt.Println(github_rest_api_CMD.Exec())
		//running github API
	// if err2 := github_rest_api_CMD.Exec(); err2 != nil{
	// 	fmt.Println("failed api req")
	// 	log.Fatal(err2)
	// }
		fmt.Println("SWAG IT WORKED")
}





	// parse_file(issues_file_name)
	
	// func parse_file(file_name string){
		
		// 	content, err := ioutil.ReadFile(file_name)
		
		
		// }

		
		// package main

// import (
//  "fmt"
//  "io/ioutil"
//  "log"
//  "strings"
// )

// func main() {
//  var path = "invoices.csv"
//  filebuffer, err := ioutil.ReadFile(path)

//  if err != nil {
//   log.Fatal(err)
//  }
//  var inputdata string = string(filebuffer)

//  rows := strings.Split(inputdata, "\n")
//  for _, row := range rows {
//   fmt.Println("row:", row)
//  }
// }