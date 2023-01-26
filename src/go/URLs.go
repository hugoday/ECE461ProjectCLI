package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
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

	// fmt.Println("[DONE]")

}
