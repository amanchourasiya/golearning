package features

import "fmt"

type Name struct {
	firstName string
	lastName  string
}

func (name *Name) getName() {
	fmt.Printf("Fullname: %s %s\n", name.firstName, name.lastName)
}

func (name *Name) setName(firstName, lastName string) {
	name.firstName = firstName
	name.lastName = firstName
}

func showFullName(name *Name) {
	fmt.Printf("Fullname: %s %s\n", name.firstName, name.lastName)
}

func MethodDemo() {
	name := Name{}
	name.setName("John", "Doe")
	name.getName()
}
