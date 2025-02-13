package channels

import (
	"fmt"
	"time"
)

// 1. Criar um canal
// 2. Iniciar uma goroutine
// 3. Enviar valores para o canal
// 4. Fechar o canal
// 5. Receber valores do canal
func ChanCounter() {
	timeStart := time.Now()

	// Cria um canal de inteiros
	ch := make(chan int)

	// Inicia uma goroutine
	go chann(ch)

	// Recebe valores do canal
	sum := 0
	for value := range ch {
		sum += value
	}

	fmt.Println("Soma dos valores:", sum)
	fmt.Println("Tempo de execução:", time.Since(timeStart).Seconds())
}

func chann(ch chan int) {
	for i := 1; i <= 100_000; i++ {
		ch <- i // Envia valores para o canal
	}
	close(ch) // Fecha o canal
}
