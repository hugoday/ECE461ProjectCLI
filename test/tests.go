package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("python", "-c", "import sys; sys.path.append('../'); from src import test; print(test.foo(1,2))")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
