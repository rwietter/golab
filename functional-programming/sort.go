package fp

import (
	"fmt"
	"sort"
)

func copySort(arr ...int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func Sort() {
	originalSlice := []int{5, 8, 2, 7, 8, 2}
	sortedSlice := copySort(originalSlice...)
	fmt.Println("Original Slice:", originalSlice) // [5 8 2 7 8 2]
	fmt.Println("Sorted Slice:", sortedSlice)     // [2 2 5 7 8 8]
}
