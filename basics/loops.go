package basics

import "fmt"

// There should be only one way of doing anything

// LoopsDemo function demonstrates different loops working
func LoopsDemo() {
	// C style for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Traditional loop\n")
	}

	// Alternative for while loop
	j := 0
	for {
		if j == 5 {
			break
		}
		j++
		fmt.Printf("Infinite for loop\n")
	}

	for j < 10 {
		fmt.Printf("while alternative for loop\n")
		j++
	}

	// Loop over iterators
	a := [5]int{1, 2, 3, 4, 5}
	for _, v := range a {
		fmt.Printf("index value:%d\n", v)
	}

}
