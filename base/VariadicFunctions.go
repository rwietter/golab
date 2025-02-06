package base

import "fmt"

func sum(nums ...int) {
	fmt.Print("...int", nums, ", ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
}

func VariadicFunctions() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
