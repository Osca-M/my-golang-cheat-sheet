package main

import "fmt"

// Employee holds information on an employee
type Employee struct {
	name   string
	salary int
	age    int
	role   string
	sex    string
}

// Employee2 a different way to define an employee
type Employee2 struct {
	name, role, sex string
	salary, age     int
}

// Building a struct with anonymous fields
type Building struct {
	string
	int
}

// Address holds Post address data
type Address struct {
	postalCode    int
	postalAddress int
	city          string
}

// Company information
type Company struct {
	name     string
	director Employee
	address  Address
}

// Student An illustration of promoted fields
type Student struct {
	name  string
	grade int
	Address
}

// A method to Employee struct, value receiver method
func (e Employee) stringRepresentation() {
	if e.sex == "Male" {
		fmt.Printf("%s is our employee, he is %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	} else if e.sex == "Female" {
		fmt.Printf("%s is our employee, she is %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	} else {
		fmt.Printf("%s is our employee, they are %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	}
}

// A pointer receiver method
func (a *Address) fullAddress() {
	fmt.Printf("Address: %d-%d, %s \n", a.postalAddress, a.postalCode, a.city)
}

// A function that uses Employee struct as a paramater
func stringRepresentation(e Employee) {
	if e.sex == "Male" {
		fmt.Printf("%s is our employee, he is %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	} else if e.sex == "Female" {
		fmt.Printf("%s is our employee, she is %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	} else {
		fmt.Printf("%s is our employee, they are %d years old, holds a %s role and earns %d amount of money \n", e.name, e.age, e.role, e.salary)
	}
}
// Name struct
type Name struct {
	Family   string
	Personal string
}
// Email struct
type Email struct {
	Kind    string
	Address string
}
// Person struct
type Person struct {
	Name  Name
	Email []Email
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}


func main() {
	osca := Employee{
		name:   "Osca Mwongera",
		salary: 700,
		age:    24,
		role:   "Backend Software Developer",
		sex:    "Male",
	}
	osca.stringRepresentation() // use of methods
	stringRepresentation(osca)  // use of functions

	lucia := Employee{
		name:   "Lucia Gacheru",
		salary: 900,
		age:    24,
		role:   "DevOps",
		sex:    "Female",
	}
	lucia.stringRepresentation()

	bot := Employee{
		name:   "Zuri Analina",
		salary: 0,
		age:    5,
		role:   "Bot",
	}
	bot.stringRepresentation()

	// Anonymous structs
	will := struct {
		name   string
		salary int
		age    int
		role   string
	}{
		name:   "Nya Ti",
		salary: 2500,
		age:    49,
		role:   "Lead",
	}
	fmt.Println("Anonymous struct", will)
	buildingA := Building{
		string: "Elodret Daima tower",
		int:    15,
	}
	fmt.Printf("%s has %d floors \n", buildingA.string, buildingA.int)

	// Nested fields

	c := Company{
		name: "Rich Oak Tech",
		director: Employee{
			name:   "Osca",
			salary: 42500,
			age:    23,
			role:   "Director",
		},
		address: Address{
			postalCode:    00100,
			postalAddress: 56454,
			city:          "Nairobi",
		},
	}
	fmt.Println("Company name:", c.name)
	fmt.Println("Director's name:", c.director.name)
	fmt.Println("Director's salary:", c.director.salary)
	fmt.Println("Director's age:", c.director.age)
	fmt.Println("Director's role:", c.director.role)
	fmt.Println("Company Postal Code:", c.address.postalCode)
	fmt.Println("Company Postal Address:", c.address.postalAddress)
	fmt.Println("Company Postal City:", c.address.city)
	s := Student{
		name:  "Victor",
		grade: 4,
		Address: Address{
			postalCode:    00100,
			postalAddress: 54128,
			city:          "Nairobi",
		},
	}
	fmt.Println("Student's name:", s.name)
	fmt.Println("Student's grade:", s.grade)
	fmt.Println("Student's postal code:", s.postalCode)       // Promoted field
	fmt.Println("Student's postal address:", s.postalAddress) // Promoted field
	fmt.Println("Student's city:", s.city)                    // Promoted field
	s.fullAddress()                                           // Accessing methods to an anonymous struct field

	// A compoiund struct
	person := Person{
		Name: Name{Family: "Mukiri", Personal: "Mwongera"},
		Email: []Email{
			Email{
				Kind: "work", Address: "m@work.com",
			},
			Email{
				Kind: "home", Address: "m@home.com",
			},
		},
	}
	fmt.Println(person.String())
}
