package fp

import "fmt"

type IntList []int

/*
A functor is a functional programming concept that represents an object or data
structure that can be mapped through a function. In other words, a functor is a
data container that implements a mapping method (often called map) that allows
you to apply a function to each element contained in the functor, producing a
new functor with the results.
*/
func (list IntList) Map(f func(int) int) IntList {
	result := make([]int, len(list))
	for index, item := range list {
		result[index] = f(item)
	}
	return result
}

func Functor() {
	numbers := IntList{1, 2, 3, 4, 5}

	// Functor is a object that implements a mapping method
	// The mapping modifies the values inside the functor
	// but does not modify the functor itself (data type and data structure)
	doubledNumbers := numbers.Map(double)
	fmt.Println("Functor:", doubledNumbers) // Output: [4 8 12 16 20]
}
