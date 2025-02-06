package fp

import "fmt"

func Lambda() {
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 4)
	fmt.Println("Lambda:", result) // Output: 7
}
