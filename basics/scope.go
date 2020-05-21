package basics

import "fmt"

func display(a int) {
	fmt.Printf("%d\n", a)
}

func ScopeTest() {
	result := 1000
	display(result) // display1

	if true {
		result := 2000
		display(result) // display2
	}
	display(result)

	if true {
		result = 3000
		display(result) // display3
	}
	display(result) // display4
}

/*
a)          b)          c)        d)
1000        1000        1000      1000
2000        2000        2000      2000
2000        1000        1000      2000
3000        3000        3000      3000
3000        1000        3000      3000
*/
