package base

import "fmt"

func AnonFunctions() {
	result := func(value string) string {
		return fmt.Sprintf("Hello %s", value)
	}("World")

	fmt.Println(result)
}
