// maps2
// Make me compile!
package main

import "fmt"

func main() {
	m := map[string]int{
		"John": 1,
		"Ana":  2,
	}
	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
