package channels

import (
	"fmt"
	"time"
)

func Chan() {
	channel := make(chan string)
	go routine("1111111", channel)
	go routine("2222222", channel)

	for {
		message, open := <-channel // receive value from channel

		// check if channel is closed or open
		if !open {
			break
		}

		fmt.Println(message)
	}

	// alternative way to receive value from channel
	for message := range channel {
		fmt.Println(message)
	}
}

func routine(str string, channel chan string) {
	for i := 0; i < 2; i++ {
		channel <- str // send value to channel
		time.Sleep(time.Second)
	}
	close(channel) // close channel
}
