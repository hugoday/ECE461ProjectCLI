package main

import (
	"fmt"
	"os/exec"
)

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

func main() {

	fmt.Println(runModule("test.foo(2,5)"))

}
