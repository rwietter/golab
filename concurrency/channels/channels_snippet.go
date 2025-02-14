package channels

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 1. Goroutines consumers (print) não precisam estar no WaitGroup, pois elas encerram automaticamente ao detectar que os canais foram fechados
func ChannelSnippet() {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	par := make(chan int)
	impar := make(chan int)

	fmt.Println("SO Threads:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(2)
	go calculate(par, impar, &wg)
	go print(par, impar, &wg)

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("Waiting...")
	wg.Wait()

	fmt.Println("After wg.Wait - Goroutines:", runtime.NumGoroutine())

	fmt.Println("Tempo de execução:", time.Since(startTime).Seconds())
}

func calculate(par, impar chan int, wg *sync.WaitGroup) {
	defer close(par)
	defer close(impar)
	defer wg.Done() // Sinaliza que a goroutine terminou

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
}

func print(par, impar chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Sinaliza que a goroutine terminou
	for par != nil || impar != nil {
		select {
		case v, ok := <-par:
			if !ok {
				par = nil // evitar que o case seja selecionado novamente
				continue
			}
			fmt.Println("Par:", v)
		case v, ok := <-impar:
			if !ok {
				impar = nil // evitar que o case seja selecionado novamente
				continue
			}
			fmt.Println("Impar:", v)
		}
	}
}
