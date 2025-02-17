# Cond

...a rendezvous point for goroutines waiting for or announcing the occurrence of an event. (um ponto de encontro para goroutines esperando ou anunciando a ocorrência de um evento.)

Const fornece uma maneira de uma goroutina dormir com eficiência até que for sinalizada para acordar e verificar sua condição. Isso é exatamente o que o tipo COND faz por nós.

```go
c := sync.NewCond(&sync.Mutex{}) // cria um novo condicionador
c.L.Lock() // bloqueia o mutex 
for conditionTrue() == false { // enquanto a condição não for verdadeira 
  c.Wait() // espera (nesse caso, dorme ou seja, libera a thread)
}
c.L.Unlock() // desbloqueia o mutex
```

## Caso de uso

> Imagine que você tenha um conjunto de goroutines que estão esperando por um evento para ocorrer. Quando o evento ocorre, todas as goroutines devem ser notificadas e acordadas para verificar a condição.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var c = sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.L.Lock() // adquire o mutex
			for i != 0 {
				c.Wait() // enquanto i for diferente de 0, dorme e libera o mutex
			}
			fmt.Println(i) // Quando acordada, automaticamente readquire o mutex e imprime o valor de i
			c.L.Unlock()
		}(i)
	}

	c.L.Lock() // adquire o mutex antes do broadcast
	fmt.Println("Broadcasting")
	c.Broadcast() // acorda todas as goroutines que estavam esperando
	c.L.Unlock() // libera o mutex para que as goroutines acordadas possam verificar suas condições

	wg.Wait()
}
```

O acontece nesse código é que temos 10 goroutines que estão esperando por um evento para ocorrer. O evento é que `i` seja igual a 0. Quando `i` é igual a 0, todas as goroutines devem ser notificadas e acordadas para verificar a condição. Quando a condição é verdadeira, a goroutine imprime o valor de `i` e libera o mutex.

O `c.Broadcast()` pode ser chamado de qualquer goroutine que precise sinalizar para as outras goroutines que estão dormentes. No exemplo do código, o broadcast está sendo feito na thread principal (main), mas poderia ser feito por qualquer outra goroutine que precise notificar as demais.

Por exemplo, poderia ser um cenário onde:

- Uma goroutine "produtora" está processando dados
- Várias goroutines "consumidoras" estão esperando esses dados
- Quando a produtora termina, ela usa c.Broadcast() para acordar todas as consumidoras

Além do `Broadcast()`, o **Cond** também oferece o método `Signal()` que acorda apenas uma única goroutine dormente, ao invés de todas. É útil quando você só precisa notificar um consumidor, por exemplo.

## Exemplo com Signal()

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
	var queue []string
	c := sync.NewCond(&sync.Mutex{})
	done := make(chan bool)

	// Consumidor: espera por itens na fila
	go func() {
		c.L.Lock()
		for len(queue) == 0 {
				c.Wait() // espera até que um item seja adicionado
		}
		fmt.Println("Consumindo:", queue[0])
		queue = queue[1:] // remove o primeiro item
		c.L.Unlock()
		done <- true
	}()

	// Produtor: adiciona um item e sinaliza
	go func() {
		time.Sleep(time.Second) // simula algum processamento
		c.L.Lock()
		queue = append(queue, "novo item")
		fmt.Println("Produzido: novo item")
		c.Signal() // acorda apenas uma goroutine (o consumidor)
		c.L.Unlock()
	}()

	<-done
}
```

Neste exemplo:
1. O consumidor espera até que haja itens na fila
2. O produtor adiciona um item e usa `Signal()` para acordar apenas uma goroutine
3. O `Signal()` é mais eficiente que `Broadcast()` neste caso, pois só precisamos acordar um consumidor


o `Signal()` **acorda apenas uma goroutine dormente de forma não determinística**. A especificação do Go não garante qual goroutine será acordada quando há múltiplas goroutines esperando - é escolhida arbitrariamente.

Por isso que o `Signal()` é mais adequado em cenários onde:

1. Você só precisa acordar uma goroutine por vez
2. Não importa qual goroutine específica será acordada
3. Todas as goroutines dormentes executam o mesmo tipo de trabalho

Um exemplo clássico é o padrão produtor-consumidor com múltiplos consumidores idênticos, onde qualquer consumidor pode processar o item produzido.

Internamente, o tempo de execução mantém uma lista do FIFO de goroutines esperando para serem sinalizados; O Signal encontra a goroutine que espera mais tempo e notifica isso, enquanto a transmissão envia um sinal para todas as goroutines que estão esperando. A transmissão é sem dúvida o mais interessante dos dois métodos, pois fornece uma maneira de se comunicar com várias goroutines de uma só vez.