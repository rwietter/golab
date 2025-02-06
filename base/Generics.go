package base

type Summable interface {
	int | float64
}

// Type constraint.

// O operador (~) permite que qualquer tipo definido pelo usuário que tenha int ou float64 como tipo subjacente seja considerado parte da interface Number.

// Por exemplo, se você tiver um tipo definido como type MyInt int, ele será aceito pela interface Number porque seu tipo subjacente é int.
type Number interface {
	~int | ~float64
}

func sumGeneric[T Summable](a, b T) T {
	return a + b
}

func compare[T comparable](a, b T) bool {
	return a == b
}

func Generics() {
	println("sumGeneric int:", sumGeneric(1, 2))
	println("sumGeneric float64:", sumGeneric(1.1, 2.2))

	println("compare int:", compare(1, 2))
	println("compare float64:", compare(1.1, 2.2))
}
