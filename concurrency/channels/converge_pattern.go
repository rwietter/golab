package channels

import "fmt"

func ConvergePattern() {
	quant := 10
	channel := converge(pares(quant), impares(quant))

	for ch := range channel {
		println(ch)
	}
}

func pares(quant int) (channel chan string) {
	channel = make(chan string)
	go func(l int) {
		defer close(channel) // ← Fecha o canal apenas quando a goroutine termina!
		for i := 1; i <= l; i++ {
			if i%2 == 0 {
				channel <- fmt.Sprintf("Par %d", i)
			}
		}
	}(quant)
	return
}

func impares(quant int) (channel chan string) {
	channel = make(chan string)
	go func(l int) {
		defer close(channel) // ← Fecha o canal apenas quando a goroutine termina!
		for i := 1; i <= l; i++ {
			if i%2 != 0 {
				channel <- fmt.Sprintf("Impar %d", i)
			}
		}
	}(quant)
	return
}

func converge(c1, c2 chan string) (channel chan string) {
	channel = make(chan string)
	go func() {
		defer close(channel) // ← Fecha o canal apenas quando a goroutine termina!
		for c1 != nil || c2 != nil {
			select {
			case v, ok := <-c1:
				if !ok {
					c1 = nil
					continue // Return causaria a saída da goroutine
				}
				channel <- v
			case v, ok := <-c2:
				if !ok {
					c2 = nil
					continue // Return causaria a saída da goroutine
				}
				channel <- v
			}
		}
	}()
	return
}
