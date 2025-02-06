// structs1
// Make me compile!
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var person = Person{
		name: "Bob",
		age:  20,
	}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
