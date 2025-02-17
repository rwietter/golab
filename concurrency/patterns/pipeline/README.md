# Pipeline Pattern

- Conecta uma série de estágios de processamento
- Cada estágio recebe dados do anterior e envia para o próximo
- Muito útil para transformações sequenciais de dados
- No exemplo: `generator -> multiply -> filter`

## Benefícios desses padrões

- Paralelização eficiente
- Controle fino sobre número de workers
- Balanceamento natural de carga
- Ideal para CPU-bound tasks

## Casos de uso comuns

- Processamento de logs, ETL, processamento de imagens