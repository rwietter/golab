package pipeline

import (
	"context"
	"fmt"
	"time"
)

// --- Pipeline Pattern ---
type PipelineStage func(context.Context, <-chan int) <-chan int

func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func multiply(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * 2:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func filter(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 { // filtra apenas números pares
				select {
				case out <- n:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func PipelinePattern() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("=== Pipeline Pattern ===")
	// Pipeline: generator -> multiply -> filter
	numbers := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	doubled := multiply(ctx, numbers)
	evens := filter(ctx, doubled)

	for n := range evens {
		fmt.Printf("Pipeline result: %d\n", n)
	}
}
