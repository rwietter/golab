# Close Channels (End-of-data signaling)

## Regras

1. Somente o escritor pode fechar o canal, não o leitor.  Se o leitor fechá-lo, o escritor entrará em pânico na próxima escrita.
2. Um escritor só pode fechar o canal se for o único proprietário.  Se houver vários escritores e um fechar o canal, os outros entrarão em pânico na próxima escrita ou tentarão fechar o canal.
3. Channels devem ser sempre fechados para evitar vazamentos (leaks).  Mas um canal não é um recurso externo.  Quando um canal não é mais utilizado, o coletor de lixo do Go irá liberar seus recursos, esteja ele fechado ou não.

- [Next](part2.md)