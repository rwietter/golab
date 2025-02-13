# Channel iteration

Por exemplo:

1. Se eu fechar o canal pela goroutine "receiver"
2. Ler acidentalmente do canal na goroutine de "sender".

Esses erros ocorrem em tempo de execução, portanto não é possível notar até executar o programa.  Seria melhor capturá-los em tempo de compilação.

Você pode se proteger desse tipo de erro definindo a direção do canal.  Os canais podem ser:

1. `chan` (**bidirecional**): para receiver (leitura) e sender (escrita) (padrão);
2. `chan<-` (**send-only**): sender apenas para escrita;
3. `<-chan` (**receive-only**): receiver apenas leitura.