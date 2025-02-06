package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channels := mux(generator("Hello"), generator("World"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channels)
	}
}

func mux(ch1, ch2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				output <- msg
			case msg := <-ch2:
				output <- msg
			}
		}
	}()
	return output
}

// func mux(channels ...<-chan string) <-chan string {
// 	channel := make(chan string)

// 	for _, c := range channels {
// 		go func(c <-chan string) {
// 			for {
// 				channel <- <-c
// 			}
// 		}(c)
// 	}

// 	return channel
// }

func generator(value string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			timer := time.Duration(rand.Intn(5))
			channel <- fmt.Sprintf("Value: %s", timer)
			time.Sleep(time.Second * timer)
		}
	}()

	return channel
}
