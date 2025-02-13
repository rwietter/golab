# Highlights

- Canais são bloqueantes.  Isso significa que, se você enviar um valor para um canal, a goroutine que está enviando ficará bloqueada até que outro valor seja recebido do canal.  E, se você receber um valor de um canal, a goroutine que está recebendo ficará bloqueada até que outro valor seja enviado para o canal.

- os canais não apenas transferem dados, mas também ajudam a sincronizar goroutines independentes.

- O canal `ch` é fechado com `defer close(ch)` após o loop que envia os dados. Isso sinaliza ao loop principal que não há mais dados para receber, evitando deadlocks.

- Uma abordagem comum ao usar goroutines é faze-las para **Producer** e **Consumer**.
  - Processamento Assíncrono (I/O);
  - Pipeline de Dados [Producer (coleta dados) → Consumer 1 (processa dados) → Consumer 2 (armazena dados)];
  - Concurrency (múltiplos consumers em paralelo para processar dados de um único producer).

- Em goroutines, precisamos sinalizar quando que uma goroutine termina, isso para travar o código seguinte ou função superior para que não termine antes que a goroutine consumer ou producer termine o que está fazendo. Quando precisa aguardar múltiplos consumers, usamos o `sync.WaitGroup`.

- Go não garante que a escrita no mapa seja segura em cenários concorrentes sem sincronização explícita. O Go não permite acesso concorrente a mapas, mesmo que seja apenas uma goroutine escrevendo e outras lendo.

- Executar `go mod tidy` para atualizar as dependências do projeto ou após criar módulos e arquivos go.mod e go.sum.