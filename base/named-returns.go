package base

import "fmt"

func operations(x, y int) (sum, minus int) {
	sum = x + y
	minus = x - y
	return
}

func NamedReturns() {
	fmt.Println(operations(1, 2))
}
