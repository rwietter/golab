# WaitGroup

Usar `time` para esperar as goroutines terminarem não é a melhor prática. Podemos usar `sync.WaitGroup` para esperar as goroutines terminarem. Por exemplo:

`WaitGroup` possui um contador dentro.  Chamar `wg.Add(1)` aumenta em um, enquanto `wg.Done()` diminui.  E `wg.Wait()` **bloqueia a goroutine** (neste caso, `main`) até que o contador chegue a **zero**.  Dessa forma, `main` espera que `say(1)` e `say(2)` terminem antes de sair.

```go
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Add a count of two, one for each goroutine.
	wg.Add(1)
	go say(&wg, 1, "Hello, how are you?")
	wg.Add(1)
	go say(&wg, 2, "Hi, I'm fine, thank you!")
	wg.Wait() // Wait for all the goroutines to finish.
}

func say(wg *sync.WaitGroup, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}

	wg.Done() // Signal that this goroutine is done.
}
```

No entanto, esta abordagem mistura lógica de negócios (digamos) com lógica de simultaneidade (wg).  Como resultado, não podemos executar facilmente, digamos, em código regular e não simultâneo.

- [Next](part3.md)