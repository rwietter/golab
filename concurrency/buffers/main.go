package main

func main() {
	// create buffered channel
	// buffer defines how many values can be stored in channel
	// Send and receive operations block only when the buffer is full
	channel := make(chan string, 2)

	// send value to channel
	channel <- "1111111"

	message := <-channel
	println(message)

	// message = <-channel
	// println(message)
}
