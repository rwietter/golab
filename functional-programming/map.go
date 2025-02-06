package fp

import "fmt"

func mapping(s []int, f func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func Mapping() {
	numbers := []int{1, 2, 3, 4, 5}

	doubledNumbers := mapping(numbers, double)

	fmt.Println("Mapping:", doubledNumbers) // Output: [2 4 6 8 10]

	// Using lambda function
	doubled := mapping(numbers, func(x int) int {
		return x * 2
	})

	fmt.Println("Lambda mapping:", doubled)
}
