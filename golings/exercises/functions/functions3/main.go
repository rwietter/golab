// functions3
// Make me compile!

package main

import "fmt"

func main() {
	const num = 10
	call_me(num)
}

func call_me(num int) {
	for n := 0; n <= num; n++ {
		fmt.Printf("Num is %d\n", n)
	}
}
