package fp

import "fmt"

func add1(args ...int) int {
	var result int
	for _, v := range args {
		result += v
	}
	return result
}

func partial(fn func(...int) int, value int) func(...int) int {
	return func(args ...int) int {
		return fn(append([]int{value}, args...)...)
	}
}

func PartialApplication() {
	sum := partial(add1, 5)

	result := sum(3, 2)

	fmt.Println("Partial Application:", result) // Output: 10 ((5) + 3 + 2)
}
