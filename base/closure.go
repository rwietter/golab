package base

func closure() func(world string) {
	hello := "Hello"

	fn := func(world string) {
		println(hello, world)
	}

	return fn
}

func Closure() {
	fn := closure()
	fn("World")
}
