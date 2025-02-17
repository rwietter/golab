# Goroutines

Goroutines não são threads do sistema operacional e não são exatamente green threads - os threads gerenciadas pelo tempo de execução de uma linguagem - são um nível mais alto de abstração conhecido como Coroutines. As coroutinas são simplesmente sub-rotinas concorrentes (funções, closures ou métodos em Go) que não são preventivos - ou seja, não podem ser interrompidos. Em vez disso, as coroutinas têm vários pontos em todo o qual permitem suspensão ou reentrada.

O runtime Go observa o comportamento de tempo de execução das goroutines e as suspende automaticamente quando elas bloqueiam e depois as retomam quando se desbloquearam. Assim, as goroutines podem ser consideradas uma classe especial de coroutina.

Coroutines, e, portanto, goroutines, são construções implicitamente concorrentes, mas a concorrência não é uma propriedade de uma coroutina: algo deve hospedar várias coroutinas simultaneamente e dar a cada um a oportunidade de executar - além, elas não seriam concomitadas!

O mecanismo de Go para hospedar Goroutines é uma implementação do que é chamado de `scheduler M:N`, o que significa que **ele mapeia os green threads para N OS threads**. Goroutines são agendados para os green threads. Quando temos mais goroutines do que green threads disponíveis, o agendador lida com a distribuição das goroutines nos threads disponíveis e garante que, quando essas goroutines ficarem bloqueadas, outras goroutines podem ser executadas.

GO segue um modelo de concorrência chamado modelo de bifurcação (fork-join). A palavra junção (fork) refere-se ao fato de que, em algum momento no futuro, esses branchs concorrentes de execução se unirão novamente.  Onde a child se junta ao parent é chamado de ponto de junção (join point).

```md
      main
       |
       |
       | ----------------- fork -----------------> child
       |                                         |
       |                                         |
       |                   <-done                |
       | <----------------- join ----------------- child
       |
       |
      main
```

Para criar um ponto de junção, você precisa sincronizar a goroutine principal e a goroutine secundária. Isso pode ser feito com canais, `sync.WaitGroup`, `sync.Mutex`, `sync.Cond`, etc.

Aqui, a goroutine principal espera que a goroutine secundária termine de executar antes de sair. Se você remover a chamada `wg.Wait()`, a goroutine principal sairá imediatamente, e a goroutine secundária não terá a chance de terminar.

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Hello, playground")
	}()

	wg.Wait()
}
```

Este exemplo bloqueará determinadamente a goroutine main até que a goroutine secundária termine de executar.

# Detalhe importante sobre closures

Em Go, as **closures (funções anônimas) capturam variáveis do escopo externo por referência, não por valor**. Isso significa que a goroutine não recebe uma cópia do valor de salutation no momento em que é criada, mas sim uma referência à variável salutation do loop. Go observa que a variável salutation é usada dentro da goroutine e, portanto, mantém uma referência a ela na stack.

```go
var wg sync.WaitGroup

for _, salutation := range []string{"hello", "greetings", "good day"} {
  wg.Add(1)
  go func() {
    defer wg.Done()
    fmt.Println(salutation) // captura a variável salutation por referência
  }()
} 
wg.Wait()

// Saída:
// good day
// good day
// good day
```

A solução para esse problema é passar a variável salutation como um argumento para a função anônima.

```go
var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {
  wg.Add(1)
  go func(salutation string) {
    defer wg.Done()
    fmt.Println(salutation)
  }(salutation)
}
wg.Wait()
```

# Green Threads e eficiência

Uma goroutine recém criada recebe alguns kilobytes, o que é quase sempre suficiente.  Quando não é, o tempo de execução cresce (e diminui) a memória para armazenar a pilha automaticamente, permitindo que muitas goroutines vivam em uma quantidade modesta de memória. A sobrecarga da CPU em média cerca de três instruções baratas por chamada de função. É prático criar centenas de milhares de goroutines no mesmo espaço de endereço. 

No entanto a troca de contexto é o problema aqui, quando algo que hospeda um processo concorrente deve salvar seu estado para mudar para executar um processo simultâneo diferente. **Se tivermos muitos processos concorrentes, podemos gastar todo o nosso contexto de tempo da CPU alternando entre eles e nunca realizar um trabalho real.** No nível do sistema operacional, com threads, isso pode ser bastante caro. O thread do sistema operacional deve salvar coisas como valores de registro, tabelas de pesquisa e mapas de memória para poder voltar para o thread atual quando chegar a hora. Em seguida, ele deve carregar as mesmas informações para o thread de entrada.

# Sincronização

Como várias goroutines podem operar no mesmo espaço de endereço, ainda precisamos nos preocupar com a sincronização.