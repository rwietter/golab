# Channels

O modelo de multi-threading do Go é baseado em CSP (Communicating Sequential Processes).  Isso significa que as goroutines se comunicam entre si usando canais.  Canais são como tubulações, que se conectam e permitem que você envie coisas (values) de uma tubulação para outra.

Os canais são bidirecionais, ou seja, você pode enviar e receber valores através deles.  Eles são seguros para concorrência, o que significa que você pode enviar e receber valores deles de várias goroutines simultaneamente.

```
┌─────────────┐    ┌─────────────┐
│ goroutine A │    │ goroutine B │
│             └────┘             │
│        X <-  chan  <- X        │
│             ┌────┐             │
│             │    │             │
└─────────────┘    └─────────────┘
```

Aqui vemos que a goroutine B envia um valor `X` para o canal `chan`, e a goroutine A recebe esse valor `X` do canal `chan`.


**IMPORTANTE: Canais são bloqueantes.  Isso significa que, se você enviar um valor para um canal, a goroutine que está enviando ficará bloqueada até que outro valor seja recebido do canal.  E, se você receber um valor de um canal, a goroutine que está recebendo ficará bloqueada até que outro valor seja enviado para o canal.**

No exemplo a seguir, Após enviar a mensagem para o canal ➊, a goroutine B é bloqueada. Somente quando a goroutine A recebe a mensagem ➌ a goroutine B continua e imprime "mensagem enviada" ➋.

```go
func main() {
    messages := make(chan string)

    go func() {
        fmt.Println("B: Sending message...")
        messages <- "ping"                    // (1)
        fmt.Println("B: Message sent!")       // (2)
    }()

    fmt.Println("A: Doing some work...")
    time.Sleep(500 * time.Millisecond)
    fmt.Println("A: Ready to receive a message...")

    <-messages                               //  (3)

    fmt.Println("A: Messege received!")
    time.Sleep(100 * time.Millisecond)
}
```

**IMPORTANTE: os canais não apenas transferem dados, mas também ajudam a sincronizar goroutines independentes.**

- [Next](part6.md)
