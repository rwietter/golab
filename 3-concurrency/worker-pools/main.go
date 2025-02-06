package main

import (
	"fmt"
	"time"
)

// Para saber quantos cores o processador tem, execute nproc --all ou lscpu
// sem concorrência (1 core) -> 64.5 segundos
// com concorrência (4 cores) -> 48.7 segundos
func main() {
	initialTime := time.Now()
	tasks, results := make(chan int, 50), make(chan int, 50)
	// fibo := fibonacci(45)
	// println(fibo)

	// 1 worker para cara core do processador
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 50; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 50; i++ {
		println(<-results)
	}
	close(results)

	endTime := time.Now()
	diff := endTime.Sub(initialTime)
	fmt.Printf("Tempo de execução: %.6f segundos\n", diff.Seconds())
}

// tasks: só recebe dados
// results: só envia dados
func worker(tasks <-chan int, results chan<- int) {
	for n := range tasks {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
