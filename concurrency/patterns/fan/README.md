# Fan-out/Fan-in Pattern

- Fan-out: distribui o trabalho entre múltiplos workers
- Fan-in: combina os resultados de volta em um único canal
- Excelente para processamento paralelo de dados
- Útil quando:
    - O processamento é independente por item
    - A ordem de processamento não importa
    - Você quer maximizar throughput

## Benefícios desses padrões

- Paralelização eficiente
- Controle fino sobre número de workers
- Balanceamento natural de carga
- Ideal para CPU-bound tasks

## Casos de uso comuns

Processamento em batch, web scraping, computação paralela