package base

import "fmt"

func SlicesCap() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Length:", len(s), "Capacity:", cap(s))                                                     // 11, 11
	fmt.Printf("Length (1:3): %d, Capacity (1:3): %d, Slice (1:3): %v\n", len(s[1:3]), cap(s[1:3]), s[1:3]) // [1, 2]
	fmt.Printf("Length (:3): %d, Capacity (:3): %d, Slice (:3): %v\n", len(s[:3]), cap(s[:3]), s[:3])       // [0, 1, 2]
	fmt.Printf("Length (5:): %d, Capacity (5:): %d, Slice (5:): %v\n", len(s[5:]), cap(s[5:]), s[5:])       // [5, 6, 7, 8, 9, 10]
	s = append(s, 11)
	fmt.Println("Length:", len(s), "Capacity:", cap(s)) // Capacidade é duplicada pela capacidade inicial
}
