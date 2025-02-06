package main

import "time"

func main() {
	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "Hello"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "World"
		}
	}()

	/*
		for {
			// chan1 is blocked until chan2 sends a message (chan1 is blocked by 2 seconds)
			message1, message2 := <-chan1, <-chan2
			println(message1 + " " + message2)

			// Expected output: 1: Hello Hello Hello Hello World - 2: Hello Hello Hello Hello World
			// Received output: 1: Hello World - 2: Hello World
		}
	*/

	// Solution: use select
	for {
		select {
		case message1 := <-chan1:
			println(message1)
		case message2 := <-chan2:
			println(message2)
		}
		// Output: Hello Hello Hello Hello World
	}

}
