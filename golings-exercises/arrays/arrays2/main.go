// arrays2
// Make me compile!

package main

import "fmt"

type Names interface{}

func main() {
	names := [4]Names{"John", "Maria", "Carl", 10}
	fmt.Println(names)
}
