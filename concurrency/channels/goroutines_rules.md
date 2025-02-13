# Goroutines Rules

1. O canal deve ser fechado pelo remetente, nunca pelo receptor.
2. O canal deve ser fechado quando o remetente não tem mais dados para enviar: `defer close(ch)` dentro da goroutine.
3. O canal nunca deve ser fechado mais de uma vez.
4. Um for loop com select deve ter uma condição de parada para evitar deadlock `for ch != nil {}`.
5. Podemos usar `sync.WaitGroup` para aguardar a goroutine finalizar, mas podemos usar um channel `done` quando temos apenas uma goroutine producer. Isso funciona, a main thread fica travada até que o producer termine. Neste caso, quem deve fechar o canal é o consumer `close(done)` e na main thread apenas esperamos receber um valor do canal `<-done`.

# 1

```go
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch) // consumer não é uma goroutine porque a main thread deve esperar o consumer terminar. Se consumer fosse uma goroutine, a main thread terminaria antes do consumer. Solução: usar channel de controle `done` ou `sync.WaitGroup`.
}

func producer(ch chan int) {
	defer close(ch) // sinaliza que não há mais dados para enviar
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
```

# 2

```go
func main() {
	ch := make(chan int)
	go func() {
		defer close(ch) // sinaliza que não há mais dados para enviar
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
}
```

# 4 e 5

```go
func main() {
	ch := make(chan int)
  done := make(chan struct{})
	go producer(ch)
	go consumer(ch, done)
  <-done // main thread fica travada até que `done` receba um valor (close(done))
}

func producer(ch chan int) {
	defer close(ch) // sinaliza que não há mais dados para enviar
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func consumer(ch chan int, done chan struct{}) {
  defer close(done) // Ao finalizar a goroutine, sinaliza que não há mais dados para receber
	for ch != nil { // Condição de parada para evitar deadlock
		select {
		case i, ok := <-ch:
			if !ok {
        ch = nil // sinaliza que não há mais dados para receber
				continue
			}
			fmt.Println(i)
		}
	}
}
```