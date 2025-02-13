# Goroutines

Goroutines são green threads. Ou seja, são threads leves, que são gerenciadas pelo Go runtime. Elas são mais baratas que as threads tradicionais, que são gerenciadas pelo sistema operacional.

A `main` é por natureza uma goroutine. Ou seja, a função `main` é executada em uma goroutine separada. E, quando a função `main` termina, o programa finaliza, mesmo que existam outras goroutines ainda em execução.

Para criar uma goroutine, basta adicionar a palavra-chave `go` antes da chamada da função. Por exemplo:

```go
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Call the function from the package
	// The function is in the same package
	// So, we can call it directly
	go say(1, "Hello, how are you?")
	go say(2, "Hi, I'm fine, thank you!")

	/*
		Precisamos adicionar isso aqui, porque a função main
		não espera as goroutines terminarem para finalizar.
		Então, se não colocarmos isso, o programa vai finalizar
		antes das goroutines terminarem.

		Por quê isso acontece?
		Porque a função main é executada em
		uma goroutine separada. Ou seja, a função main
		é uma goroutine. E, quando a função main termina,
		o programa finaliza, mesmo que existam outras goroutines
		ainda em execução.
	*/
	time.Sleep(2 * time.Second)
}

func say(id int, phrase string) {
    for _, word := range strings.Fields(phrase) {
        fmt.Printf("Worker #%d says: %s...\n", id, word)
        dur := time.Duration(rand.Intn(100)) * time.Millisecond
        time.Sleep(dur)
    }
}
```

- [Next](part2.md)
