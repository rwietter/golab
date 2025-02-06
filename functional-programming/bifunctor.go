package fp

import (
	"fmt"
	"strings"
)

// Bifunctor é uma estrutura que opera em dois tipos diferentes.
type Bif struct {
	Value1 int
	Value2 string
}

// MapInt aplica uma função a Value1.
func (b Bif) MapInt(f func(int) int) Bif {
	return Bif{Value1: f(b.Value1), Value2: b.Value2}
}

// MapString aplica uma função a Value2.
func (b Bif) MapString(f func(string) string) Bif {
	return Bif{Value1: b.Value1, Value2: f(b.Value2)}
}

func Bifunctor() {
	bif := Bif{Value1: 42, Value2: "hello"}

	double := func(x int) int {
		return x * 2
	}

	uppercase := func(s string) string {
		return strings.ToUpper(s)
	}

	// Aplicação de mapeamento usando MapInt e MapString
	result := bif.MapInt(double).MapString(uppercase)

	fmt.Println(result) // Output: {84 HELLO}
}
