package main

//$go install github.com/go-python/gopy@latest
import (
	"fmt"
	"os/exec"
	// "github.com/go-python/gopy"
)

func main() {
	// // Initialize the Python interpreter
	// if err := gopy.Init(); err != nil {
	// 	fmt.Println("Error initializing Python interpreter:", err)
	// 	return
	// }
	// defer gopy.Finalize()

	// // Import the Python module
	// module, err := gopy.ImportModule("../src/test.py")
	// if err != nil {
	// 	fmt.Println("Error importing Python module:", err)
	// 	return
	// }

	// // Call the Python function
	// result, err := module.Call("foo", 1, 2)
	// if err != nil {
	// 	fmt.Println("Error calling Python function:", err)
	// 	return
	// }

	// // Print the result
	// fmt.Println("Result:", result)
	cmd := exec.Command("python", "-c", "import sys; sys.path.append('../'); from src import test; print(test.foo(1,2))")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
