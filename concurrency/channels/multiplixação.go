package channels

import (
	"fmt"
	"time"
)

func Multiplex() {
	// Cria dois canais
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Goroutine para enviar valores para ch1
	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch1)
	}()

	// Goroutine para enviar valores para ch2
	go func() {
		for i := 10; i <= 12; i++ {
			ch2 <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch2)
	}()

	// Usando select para receber valores de ambos os canais
	for {
		select {
		case value, ok := <-ch1:
			if !ok {
				ch1 = nil // Desativa este caso quando o canal é fechado
			} else {
				fmt.Println("Recebido de ch1:", value)
			}
		case value, ok := <-ch2:
			if !ok {
				ch2 = nil // Desativa este caso quando o canal é fechado
			} else {
				fmt.Println("Recebido de ch2:", value)
			}
		}

		// Se ambos os canais estiverem fechados, saia do loop
		if ch1 == nil && ch2 == nil {
			break
		}
	}

	fmt.Println("Fim!")
}
