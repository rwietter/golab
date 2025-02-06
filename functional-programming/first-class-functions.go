package fp

import "fmt"

/*
 * First-class functions are functions that can be assigned to a variable,
 * passed as an argument to another function and returned from another function.
 */
func processNumbers(numbers []int, cb func(int) int) []int {
	result := []int{}
	for _, num := range numbers {
		result = append(result, cb(num))
	}
	return result
}

func double(x int) int {
	return x * 2
}

func FirstClassFunction() {
	numbers := []int{1, 2, 3, 4, 5} // number slice

	doubledNumbers := processNumbers(numbers, double)

	fmt.Println("First-class functions:", doubledNumbers)
}
