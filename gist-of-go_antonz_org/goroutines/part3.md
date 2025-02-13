# defer

No Go, é comum separar a lógica de simultaneidade da lógica de negócios.  Isso geralmente é feito com funções separadas.  Em casos simples como o nosso, até funções anônimas servirão:

Aqui usamos **Defer**. O `defer wg.Done()` garante que a goroutine **diminua o contador** antes de sair, mesmo que entre em panic.

`say` em si não sabe nada sobre simultaneidade, isso mantém a lógica de negócios limpa. E `main` é responsável por lidar com a simultaneidade.

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
	wg.Add(2) // Add a count of two, one for each goroutine.

	go func() {
		defer wg.Done() // Signal that this goroutine is done.
		say(1, "Hello, how are you?")
	}()

	go func() {
		defer wg.Done() // Signal that this goroutine is done.
		go say(2, "Hi, I'm fine, thank you!")
	}()

	wg.Wait() // Wait for all goroutines to finish.
}

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}
```

- [Next](part4.md)