# Buffers

Canais de buffer são canais que aceitam um número limitado de valores antes de bloquear. Por exemplo, o canal `make(chan int)` é um canal sem buffer, enquanto o canal `make(chan int, 2)` é um canal com buffer de tamanho 2.

Isso significa que o canal sem buffer aceita apenas uma gravação (de qualquer valor) antes de bloquear, enquanto o canal com buffer aceita duas gravações antes de bloquear. Seja dois envios ou dois recebimentos.

```go
package main

import "fmt"

func main() {

    // Aqui nós criamos um canal com buffer de tamanho 2.
    // Isso significa que o canal aceitará até 2 envios
    // sem que o receptor esteja pronto para receber.
    messages := make(chan string, 2)

    // Como o canal tem buffer de tamanho 2, podemos enviar
    // esses valores para o canal sem bloquear.
    messages <- "buffered"
    messages <- "channel"

    // Então podemos receber esses dois valores como
    // de costume.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

Sem buffer, precisamos contornar o canal para que ele aceite um envio, bloqueie, aceite um recebimento, bloqueie, etc. Com buffer, podemos enviar até o buffer ser preenchido.

```go
package main

import "fmt"

func main() {

    // Aqui nós criamos um canal sem buffer, que bloqueará
    // após receber um valor sem que um receptor esteja pronto
    // para receber esse valor.
    messages := make(chan string)

    // Como não há buffer, precisamos de um receptor pronto
    // antes de enviar um valor para o canal. Aqui estamos
    // enviando "ping" para o canal `messages` que criamos
    // acima, de uma nova goroutine.
    go func() {
      messages <- "ping"
    }()

    // Aqui estamos recebendo o valor "ping" que enviamos
    // acima e imprimindo-o.
    msg := <-messages
    fmt.Println(msg)
}
```
