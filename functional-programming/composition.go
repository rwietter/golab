package fp

import "fmt"

func f(x int) int {
	return x * 2
}

func g(x int) int {
	return x + 1
}

func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func Compose() {
	result := compose(f, g)(5)

	// Output: 12
	fmt.Println("Compose (f(g(5)) = f(5+1) = f(6) = 6*2 = 12):", result)
}
