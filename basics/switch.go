package basics

import "fmt"

func SwitchTest(v interface{}) {
	// x := v.(type)
	// Switch cases need not be constants
	// fallthrough
	switch v.(type) {
	case string:
		fmt.Println("Value received is of type string")
	case int:
		fmt.Println("Value received is of type int")
	case float64:
		fmt.Println("Value received is of type float32")
	default:
		fmt.Println("Value type unknown")
	}

}
