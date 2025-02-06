package base

import (
	"fmt"
	"time"
)

func fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func Recursion() {
	start := time.Now()
	fibo := fibonacci(uint64(50))
	println(fibo)

	end := time.Now()

	diff := end.Sub(start)
	fmt.Println("Tempo de execução: ", diff.Seconds())
}
