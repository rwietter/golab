package base

import (
	"errors"
	"fmt"
)

func Kinds() {
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java) // 0 false false false

	// Short variable declarations
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 3
	fmt.Println(k) // 3

	// Variables declared without an explicit initial value are given their zero value.
	// The zero value is:
	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
	var j int
	var l float64
	var m bool
	var n string
	fmt.Printf("%v %v %v %q\n", j, l, m, n) // 0 0 false ""

	// The expression T(v) converts the value v to the type T.
	// Some numeric conversions:
	// var i int = 42
	// var f float64 = float64(i)
	// var u uint = uint(f)
	// Or, put more simply:
	// i := 42
	// f := float64(i)
	// u := uint(f)

	// Type inference
	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	// When the right hand side of the declaration is typed, the new variable is of that same type:
	// var i int
	// j := i // j is an int
	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
	// i := 42           // int
	// f := 3.142        // float64
	// g := 0.867 + 0.5i // complex128
	// Try changing the initial value of v in the example code and observe how its type is affected.

	// Constants

	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World) // Hello 世界
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth) // Go rules? true

	// Numeric Constants
	// Numeric constants are high-precision values.
	// An untyped constant takes the type needed by its context.
	// Try printing needInt(Big) too.
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(needInt(Small)) // 42
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big)) // 1.2676506002282295e+29

	// Errors
	var myError error = errors.New("My error string")
	fmt.Println(myError) // My error string
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
