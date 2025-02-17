package mutex

import (
	"fmt"
	"sync"
)

func Mutex() {
	var count int
	var lock sync.Mutex

	var arithmetic sync.WaitGroup

	// Increment
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func(count *int, lock *sync.Mutex) {
			defer arithmetic.Done()
			increment(&count, &lock)
		}(&count, &lock)
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func(count *int, lock *sync.Mutex) {
			defer arithmetic.Done()
			decrement(&count, &lock)
		}(&count, &lock)
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

func increment(count **int, lock **sync.Mutex) {
	(*lock).Lock()
	defer (*lock).Unlock()
	**count++
	fmt.Printf("Incrementing: %d\n", **count)
}

func decrement(count **int, lock **sync.Mutex) {
	(*lock).Lock()
	defer (*lock).Unlock()
	**count--
	fmt.Printf("Decrementing: %d\n", **count)
}
