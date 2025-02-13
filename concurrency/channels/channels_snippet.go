package channels

import (
	"fmt"
	"sync"
	"time"
)

// 1. Goroutines consumers (print) não precisam estar no WaitGroup, pois elas encerram automaticamente ao detectar que os canais foram fechados
func ChannelSnippet() {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	par := make(chan int)
	impar := make(chan int)

	wg.Add(1)
	go calculate(par, impar, &wg)
	go print(par, impar)

	wg.Wait()

	fmt.Println("Tempo de execução:", time.Since(startTime).Seconds())
}

func calculate(par, impar chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Sinaliza que a goroutine terminou

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func print(par, impar chan int) {
	for {
		select {
		case v, ok := <-par:
			if !ok {
				par = nil // evitar que o case seja selecionado novamente
				return
			}
			fmt.Println("Par:", v)
		case v, ok := <-impar:
			if !ok {
				impar = nil // evitar que o case seja selecionado novamente
				return
			}
			fmt.Println("Impar:", v)
		}
	}
}
