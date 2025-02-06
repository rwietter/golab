package main

import (
	"fmt"
	"time"
)

func main() {
	dayOfWeek := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	fmt.Printf("Today is %s\n", dayOfWeek[int(time.Now().Weekday())])
}
