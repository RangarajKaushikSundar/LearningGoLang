package main

import "fmt"

// Aggregate Data Type: Simple struct
type person struct {
	firstName string
	lastName  string
	age       int
}

// Embedded structs
type awesomePerson struct {
	person
	hasACoolName bool
}

func main() {

	// Usage of simple structs
	fmt.Println("Simple Structs")
	james := person{
		firstName: "James",
		lastName:  "Bond",
		age:       37,
	}

	fmt.Println(james.firstName, james.lastName)

	fmt.Println()
	// Usage of embedded structs
	fmt.Println("Embedded Structs")
	boaty := awesomePerson{
		person: person{
			firstName: "Boaty",
			lastName:  "Mc Boatface",
			age:       1000,
		},
		hasACoolName: true,
	}
	// In case there is collision, you can choose to include the inner type "person"
	// Normally, all the embedded structs are "promoted" to the outer type "awesomePerson"
	fmt.Println(boaty.firstName, boaty.person.lastName, boaty.hasACoolName)

	fmt.Println()
	// Usage of anonymous structs
	fmt.Println("Anonymous Structs")
	pug := struct {
		breed string
		name  string
	}{
		breed: "pug",
		name:  "mc funny face",
	}
	fmt.Println(pug)
}
