package fp

import "fmt"

func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func Closure() {
	increment := closure()

	fmt.Println("Closure:", increment(), increment(), increment()) // 1 2 3
}

// ------------------------------------------------------------
// Another example with lambda functions
// ------------------------------------------------------------
func inc() {
	x := 5
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment()) // Output: 6
}
