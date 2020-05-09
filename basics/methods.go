package basics

type FullName struct {
	firstName string
	lastName  string
}

func (name *FullName) SetFirstName() {
	name.firstName = "aman"
}

func (name *FullName) SetLastName() {
	name.lastName = "Chourasiya"
}
