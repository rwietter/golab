# Mutex

Um mutex é uma abreviação de "mutual exclusion" e é uma maneira de garantir que duas threads não estejam em uma seção crítica ao mesmo tempo. Isso é feito através de um mecanismo de bloqueio. Se uma thread tentar bloquear um mutex que já está bloqueado, ela será bloqueada até que o mutex seja desbloqueado.


As seções críticas são podem corresponder a um gargalo em seu programa. É um pouco caro entrar e sair de uma seção crítica e, portanto, geralmente as pessoas tentam minimizar o tempo gasto em seções críticas.

Uma estratégia para isso é reduzir a seção transversal da seção crítica. Pode haver memória que precisa ser compartilhada entre vários processos simultâneos, mas talvez nem todos esses processos leiam e escrevam para essa memória. Se for esse o caso, você pode aproveitar um tipo diferente de mutex: sync.rwmutex.

