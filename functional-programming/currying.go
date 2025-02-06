package fp

import "fmt"

func add(x, y int) int {
	return x + y
}

func curry(x int) func(int) int {
	return func(y int) int {
		return add(x, y)
	}
}

func Currying() {
	result := curry(5)(10)

	fmt.Println("Currying:", result) // 15
}
