// structs2
// Make me compile!
package main

import "fmt"

type ContactDetails struct {
	phone string
}

type Person struct {
	ContactDetails
	name string
	age  int
}

func main() {
	person := Person{
		name: "John",
		age:  32,
		ContactDetails: ContactDetails{
			phone: "123-456-7890",
		},
	}

	fmt.Printf(
		"%s is %d years old and his phone is %s\n",
		person.name,
		person.age,
		person.phone,
	)
}
