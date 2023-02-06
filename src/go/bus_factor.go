package main

import (
	"os"
	"fmt"
	"log"
)

func main(){

	// repo := "kubernetes"
	repo_dir := "./"
	print(repo_dir)
	err := os.Chdir(repo_dir)
	if err != nil{
		print("Did not find cloned repo")
		log.Fatal(err)
	}
	wd, _ := os.Getwd()
	print(wd)


}


func print (in string){
	fmt.Println(in)
}