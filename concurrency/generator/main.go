package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		channel := generator(fmt.Sprintf("%d", i))
		fmt.Println(<-channel)
	}
}

// cria uma goroutine que retorna um channel
// criar abstrações para trabalhar com concorrência
func generator(value string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", value)
			time.Sleep(time.Second * 1)
		}
	}()

	return channel
}
