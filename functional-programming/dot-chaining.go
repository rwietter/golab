package fp

import "fmt"

type Ints []int

func (ints Ints) Map(f func(int) int) Ints {
	result := make(Ints, len(ints))
	for i, v := range ints {
		result[i] = f(v)
	}
	return result
}

func (ints Ints) Filter(f func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func DotChaining() {
	numbers := Ints{1, 2, 3, 4, 5}

	result := numbers.
		Map(func(x int) int { return x * 2 }).
		Filter(func(x int) bool { return x > 5 })

	fmt.Println("Dot-chaining:", result) // Output: [6 8 10]
}
