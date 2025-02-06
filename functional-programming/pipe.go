package fp

import "fmt"

func doubleX(x int) int {
	return x * 2
}

func add10(x int) int {
	return x + 10
}

func pipe(input int, fps ...func(int) int) int {
	result := input
	for _, f := range fps {
		result = f(result)
	}
	return result
}

func Pipe() {
	pipeline := pipe(5, doubleX, add10)

	fmt.Println("Pipeline:", pipeline)
}
