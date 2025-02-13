package main

import (
	"fmt"
	"strings"
)

func main() {
	// EndOfDataSignaling()

	ch := make(chan int)
	sendOnly(ch)
	receiveOnly(ch)
}

// PART 1: Close Channels (End-of-data signaling)
func EndOfDataSignaling() {
	str := "one,two,,four"

	in := make(chan string)
	go func() {
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
		close(in)
	}()

	for {
		word, isOpen := <-in
		if !isOpen {
			break
		}
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}

	// fatal error: all goroutines are asleep - deadlock!
	/* for word := range in {
		fmt.Printf("%s", word)
	} */

	// fatal error: all goroutines are asleep - deadlock!
	/* for {
		word := <-in
		if word != "" {
			fmt.Printf("%s ", word)
		}
	} */
}

// PART 2: Directional channels
func sendOnly(stream chan<- int) {
	go func() {
		stream <- 1
		close(stream)
	}()
}

func receiveOnly(stream <-chan int) {
	go func() {
		for s := range stream {
			fmt.Println("Hell", s)
		}
	}()
}
