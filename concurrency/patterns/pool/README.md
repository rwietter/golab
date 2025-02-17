## Explicação do `worker_pool_pattern.go`

1. Cria um novo pool que pode rodar jobs em 3 goroutines, mas pode rodar 10 jobs por vez nessas 3 goroutines.
2. O start é acionado e cria uma goroutine para cada worker e aguarda para fechar os canais quando a goroutine tiver terminado.
2. Na goroutine com for na função WorkerPoolPattern, tá submetendo 20 jobs. Se buffer é de 10 e tenho 20 jobs, os outros 10 jobs são rejeitados (ver seção [Retry](#3))
3. Após o submit, o job é enviado o canal de jobs do WorkerPool ou rejeitado se o buffer do canal tiver cheio
4. O Worker é um loop infinito com select que recebe os jobs enviados pelo submit e processa eles, então o resultado é enviado para o canal de results do WorkerPool se o contexto não tiver sido cancelado.
5. Após todo os jobs (20? ou 10?) terem sido processado, ele volta pra função WorkerPoolPattern e sai do forloop e chega no final da goroutine e da stop no pool e fecha o canal de jobs e sinaliza que o WorkerPool foi fechado.
6. Volta pra função WorkerPoolPattern e faz um range no canal de WorkerPool results para ver os dados processados.

---

## Estrutura do Worker Pool

1. O pool é criado com

  - 3 workers (goroutines) que processam jobs concorrentemente
  - Buffer de 10 jobs (capacidade máxima do canal de jobs)

2. Fluxo de Execução:

  1. Inicialização:

  - Start() cria 3 goroutines (workers)
  - Cada worker roda em loop infinito esperando jobs
  - Uma goroutine extra monitora o término dos workers

  2. Submissão de Jobs:

  - Tenta enviar 20 jobs para o pool
  - Como o buffer é 10, o que acontece é:
    - Os primeiros 10 jobs entram no buffer
    - Para os próximos 10 jobs:
      - Se o buffer estiver cheio: job é rejeitado (retorna erro "sistema sobrecarregado")
      - Se houver espaço (porque workers processaram alguns jobs): job é aceito
    > Ou seja: não existe "fila de espera" - jobs são rejeitados se o buffer estiver cheio

  3. Processamento:

  - Os 3 workers pegam jobs do canal concorrentemente
  - Cada worker:
    - Pega um job do canal
    - Processa o job (simulado com delay de 100ms)
    - Envia resultado para o canal de resultados
    - Volta a pegar próximo job disponível

  4. Finalização:

  - Após tentar submeter os 20 jobs, chama Stop()
  - Stop() fecha o canal de jobs
  - Workers percebem canal fechado e terminam
  - Último worker fecha o canal de resultados

  5. Coleta de Resultados:

  - O range no canal de resultados recebe todos os jobs processados com sucesso
  - Podem ser menos que 20 jobs se alguns foram rejeitados por buffer cheio

---

## Retry

Existem várias estratégias comuns na comunidade Go para lidar com backpressure (sobrecarga). Aqui estão as principais abordagens:

**1. Retry com Exponential Backoff**

Ao receber rejeição, tentar novamente após um delay, **cada tentativa aumenta o tempo de espera exponencialmente**.

```go
for tentativas := 0; tentativas < maxTentativas; tentativas++ {
    if err := pool.Submit(job); err == nil {
        break
    }
    time.Sleep(time.Second * math.Pow(2, float64(tentativas)))
}
```

**2. Queue Externa**

- Usar uma fila persistente (Redis, RabbitMQ, etc)
- Jobs rejeitados vão para a fila
- Um consumer tenta resubmeter periodicamente
- Muito comum em sistemas distribuídos

**3. Circuit Breaker**

- Detecta quando sistema está sobrecarregado
- Para de aceitar novos jobs temporariamente
- Retoma gradualmente quando sistema normaliza
- Implementado por libs como gobreaker

**4. Rate Limiting**

- Limita número de submissões por tempo
- Usa tokens/semáforos para controlar fluxo
- Libs populares: golang.org/x/time/rate

**5. Buffer Dinâmico**

- Ajusta tamanho do buffer baseado em métricas
- Aumenta/diminui workers conforme demanda
- Requer monitoramento cuidadoso

**6. Priorização**

- Jobs críticos têm canal próprio
- Jobs menos importantes podem ser descartados
- Implementa diferentes níveis de QoS

**6. A escolha depende dos requisitos**

- Jobs críticos que não podem ser perdidos → Queue Externa
- Sistema em tempo real → Rate Limiting
- Alta disponibilidade → Circuit Breaker
- Jobs independentes → Retry simples
- Sistema distribuído → Combinação de estratégias

**7. O importante é definir claramente**

- Qual impacto de perder jobs?
- Qual latência máxima aceitável?
- Quanto recurso disponível?