package basics

import "fmt"

var (
	name string
	IsOk bool
)

const (
	pi = 3.14
)

const (
	a = iota
	b
	c
)

func VariablesDemo() {
	// Simple variable declaration and initialization
	var i int
	i = 10

	// Short hand variable declaration and initialization
	j := 20

	//j = true

	fmt.Printf("i:%d j:%d %d\n", i, j, c)

}

type Name struct {
	FirstName string
	LastName  string
}

// Declaring Arrays

var arr = [3]int{1, 2, 3}

func ArrayDemo() {
	fmt.Printf("Array element %d\n", arr[2])
}

// Slice of arrays
var slice []int = []int{}

func SliceDemo() {
	slice = append(slice, 10)
	fmt.Printf("Slice element %d\n", slice[0])
}

//Maps demo
func MapsDemo() {
	//m := map[int]string
}
