package features

import "fmt"

type Name struct {
	firstName string
	lastName  string
}

func (name *Name) getFullName() {
	fmt.Printf("Fullname: %s %s\n", name.firstName, name.lastName)
}

func showFullName(name *Name) {
	fmt.Printf("Fullname: %s %s\n", name.firstName, name.lastName)
}
