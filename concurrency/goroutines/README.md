# Goroutines

Goroutines são threads leves gerenciadas pelo tempo de execução do Go. 

`go f(x, y, z)` 

Inicia uma nova goroutine rodando f(x, y, z). A avaliação de f, x, y e z acontece na goroutine atual e a execução de f acontece na nova goroutine.

Goroutines rodam no mesmo espaço de endereço, então o acesso a memória compartilhada deve ser sincronizado.

O pacote de sincronização fornece primitivas úteis, embora você não precise muito delas em Go, pois existem outras primitivas.
