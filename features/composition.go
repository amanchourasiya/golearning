package features

type Employee struct {
	name Name
	//Name
	age int
}

func CompositionDemo() {
	emp := Employee{}

	emp.age = 25
	emp.name.firstName = "John"
	emp.name.lastName = "Doe"
	//emp.Name = Name{firstName: "John", lastName: "Doe"}
	//emp.getName()

}
