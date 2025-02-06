package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Number of goroutines to wait for

	go func() {
		result := performTask("Task 1")
		fmt.Println(result)
		waitGroup.Done() // Decrement the counter by 1
	}()

	go func() {
		result := performTask("Task 2")
		fmt.Println(result)
		waitGroup.Done() // Decrement the counter by 1
	}()

	waitGroup.Wait() // Wait until the counter is 0. Then, all goroutines have finished

	endTime := time.Now()
	fmt.Println("Total time taken:", endTime.Sub(startTime))
}

func performTask(taskName string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Completed %s", taskName)
}
