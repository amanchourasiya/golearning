package basics

import "fmt"

// FunctionDemo1 adds two function
func FunctionDemo1(i int, j int) int {
	return i + j
}

func FunctionDemo2(i, j int) (int, bool) {
	return i + j, true

}

// unexportes function
func unexportedFunction() {
	fmt.Println("Unexported function")
}
