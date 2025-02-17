# RWMutex

O sync.rwmutex é conceitualmente a mesma coisa que um mutex: protege o acesso à memória; No entanto, o RWMutex oferece um pouco mais de controle sobre a memória.

Pode solicitar um bloqueio para a leitura; nesse caso, você terá acesso a menos que o bloqueio esteja sendo mantido para escrever. Isso significa que um número arbitrário de leitores pode segurar uma trava do leitor, desde que nada mais esteja segurando uma trava do escritor.