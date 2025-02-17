package fan

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func generatorNums(ctx context.Context, nums ...int) <-chan int {
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

// --- Fan-out/Fan-in Pattern ---
func fanOut(ctx context.Context, in <-chan int, numWorkers int) []<-chan int {
	outputs := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		outputs[i] = someProcessWorker(ctx, in, i)
	}
	return outputs
}

func fanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedChan := make(chan int)

	// Função para encaminhar valores de um canal para o canal multiplexado
	forward := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case multiplexedChan <- i:
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go forward(c)
	}

	// Fecha o canal multiplexado quando todos os canais de entrada foram processados
	go func() {
		wg.Wait()
		close(multiplexedChan)
	}()

	return multiplexedChan
}

func someProcessWorker(ctx context.Context, in <-chan int, workerID int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			// Simula processamento com tempos diferentes por worker
			time.Sleep(time.Duration(workerID+1) * 100 * time.Millisecond)
			select {
			case out <- n * n: // cada worker calcula o quadrado
				fmt.Printf("Worker %d processou: %d\n", workerID, n)
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func FanOutFanInPattern() {
	fmt.Println("\n=== Fan-out/Fan-in Pattern ===")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fan-out/Fan-in
	input := generatorNums(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Fan-out para 3 workers
	workerChannels := fanOut(ctx, input, 3)

	// Fan-in combina resultados
	results := fanIn(ctx, workerChannels...)

	// Processa resultados finais
	for result := range results {
		fmt.Printf("Fan-in result: %d\n", result)
	}
}
