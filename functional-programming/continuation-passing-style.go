package fp

import (
	"fmt"
	"sync"
	"time"
)

func asyncAdd(a, b int, cont func(int)) {
	go func() {
		result := a + b
		time.Sleep(2 * time.Second)
		cont(result)
	}()
}

func ContinuationPassingStyle() {
	var wg sync.WaitGroup
	wg.Add(1)

	asyncAdd(3, 4, func(result int) {
		fmt.Println(" -> Async Add:", result)
		wg.Done()
	})

	fmt.Print("Running...")
	wg.Wait()
}
