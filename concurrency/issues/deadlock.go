package issues

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	mu    sync.Mutex
	value int
}

/*
o problema ocorre por causa do sleep que representa uma espera para que ambas as goroutines possam executar simultaneamente antes que uma termine e evite o deadlock, certo ? Então a ideia aqui é que:

go printSum(&a, &b):
Lock a memória de b e dorme 2 segundos

go printSum(&b, &a):
Lock a memória de b e dorme dois segundos

ambas as goroutines acordam e go printSum(&a, &b) tenta bloquear b e go printSum(&b, &a) tenta bloquear a mas ambas já estão bloqueadas nas diferentes threads, então o programa fica em estado de espera porque não pode prosseguir sem conseguir bloquear a ou b.
*/
func Deadlock() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *Value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b Value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

/*
Este é um exemplo clássico de deadlock conhecido como "Circular Wait" ou "Espera Circular":

1. Setup Inicial:

// Duas goroutines tentando acessar dois recursos
go printSum(&a, &b)  // Goroutine 1
go printSum(&b, &a)  // Goroutine 2

2. Sequência que causa o Deadlock:
Tempo   Goroutine 1         Goroutine 2
─────────────────────────────────────────
t=0     Lock(a)             Lock(b)
        ✓ obtém lock de a   ✓ obtém lock de b

t=2s    Tenta Lock(b)       Tenta Lock(a)
        ⚠️ bloqueada        ⚠️ bloqueada
        (esperando b)       (esperando a)


3. Por que é um Deadlock?

Goroutine 1:

- Tem o lock de a
- Está esperando o lock de b
- Nunca vai liberar a até conseguir b

Goroutine 2:

- Tem o lock de b
- Está esperando o lock de a
- Nunca vai liberar b até conseguir a

4. Visualização do Ciclo:

Goroutine 1 ── tem ─→ Lock(a) ←─ quer ── Goroutine 2
     │                                        │
     └── quer ─→ Lock(b) ←─ tem ──────────────┘

5. Por que o defer não ajuda?

```
v1.mu.Lock()
defer v1.mu.Unlock()  // Só executa quando a função retornar
time.Sleep(2 * time.Second)

v2.mu.Lock()  // Bloqueia aqui, função nunca retorna
defer v2.mu.Unlock()
```

- O defer só executa quando a função retorna
- A função nunca retorna porque está bloqueada
- Os locks nunca são liberados

6. Soluções Possíveis:

a) Ordenação consistente de locks:

func printSum(v1, v2 *Value) {
	// Sempre pegar locks na mesma ordem
	if v1 < v2 {  // Compara endereços de memória
		v1.mu.Lock()
		v2.mu.Lock()
	} else {
		v2.mu.Lock()
		v1.mu.Lock()
	}
	defer v2.mu.Unlock()
	defer v1.mu.Unlock()

	fmt.Printf("sum=%v\n", v1.value+v2.value)
}

b) Usar um único mutex para ambos os valores:

```
type Values struct {
	mu     sync.Mutex
	value1 int
	value2 int
}
```

c) Usar estruturas mais avançadas:
	- Channels para comunicação
	- atomic.Value para operações atômicas
	- sync.Map para acesso concorrente

Este tipo de problema é comum em sistemas concorrentes e distribuídos, onde múltiplos processos competem por recursos compartilhados. É importante sempre ter uma estratégia clara para aquisição de locks para evitar deadlocks.

*/
