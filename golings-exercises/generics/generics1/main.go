// generics1
// Make me compile!

package main

import "fmt"

type V interface{}

func main() {

	print("Hello, World!")
	print(42)
}

func print[K V](value K) {
	fmt.Println(value)
}
