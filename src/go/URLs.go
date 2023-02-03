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

type URLnode struct {
	url  string
	next *URLnode
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

func cloneRepo(url string) string {
	s := subprocess.New("git clone " + url + " src/repos/rnd")
	if err := s.Exec(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return ("ERROR")
	}
	return string("SUCCESS")
}

func traverseList(next URLnode) {
	for {
		fmt.Println(next.url)
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

	head := URLnode{}
	head.url = "HEAD"
	prev := &head

	for scanner.Scan() {
		new := URLnode{}
		new.url = scanner.Text()
		prev.next = &new
		prev = &new
	}

	s := subprocess.New("../../../test1.sh")
	fmt.Println(s.Exec())

	// traverseList(head)

	// fmt.Println()
	// fmt.Println("Cloning...")
	// clearRepoFolder()

	// url := head.next
	// for url != nil {
	// 	fmt.Println(cloneRepo(url.url))
	// 	url = url.next

	// 	clearRepoFolder()
	// }

	fmt.Println("[DONE]")

}

func getCorrectness() {

	regex, _ := regexp.Compile("\"total_count\": [0-9]+") //Regex for finding count of issues in input file
	num_regex, _ := regexp.Compile("[0-9]+")              //Regex for parsing count into only integer

	//closed issues
	data_closed, err1 := os.ReadFile("closed.txt")
	if err1 != nil {
		fmt.Println("Did not find closed issues file from api")
		log.Fatal(err1)
	}
	closed_count := regex.FindString(string(data_closed))
	closed_count = num_regex.FindString(closed_count)

	//open issues
	data_open, err := os.ReadFile("open.txt")
	if err != nil {
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
