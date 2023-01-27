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
)

type URLnode struct {
	url  string
	next *URLnode
}

type URLhead struct {
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
	cmd := exec.Command("git", "clone ", url)
	// fmt.Println(cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return ("ERROR")
	}
	return string("SUCCESS")
}

func traverseList(next URLnode) {
	for true {
		fmt.Println(next.url)
		if next.next == nil {
			break
		}
		next = *next.next
	}
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

	// fmt.Println("Traversing linked list...")

	traverseList(head)

	s := subprocess.New("git clone https://github.com/hugoday/resume")

	if err := s.Exec(); err != nil {
		log.Fatal(err)
	}

	url := head.next
	for url != nil {
		fmt.Println(cloneRepo(url.url))
		url = url.next
	}

	// fmt.Println("[DONE]")

}
